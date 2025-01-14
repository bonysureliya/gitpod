/**
 * Copyright (c) 2022 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { injectable, inject } from "inversify";
import * as express from "express";
import { SessionHandlerProvider } from "../session-handler";
import { Authenticator } from "../auth/authenticator";
import { UserService } from "../user/user-service";
import { OIDCCreateSessionPayload } from "./iam-oidc-create-session-payload";
import { log } from "@gitpod/gitpod-protocol/lib/util/logging";
import { User } from "@gitpod/gitpod-protocol";
import { BUILTIN_INSTLLATION_ADMIN_USER_ID, TeamDB } from "@gitpod/gitpod-db/lib";

@injectable()
export class IamSessionApp {
    @inject(SessionHandlerProvider)
    protected readonly sessionHandlerProvider: SessionHandlerProvider;
    @inject(Authenticator)
    protected readonly authenticator: Authenticator;
    @inject(UserService)
    protected readonly userService: UserService;
    @inject(TeamDB)
    protected readonly teamDb: TeamDB;

    public getMiddlewares() {
        return [express.json(), this.sessionHandlerProvider.sessionHandler, ...this.authenticator.initHandlers];
    }

    public create(): express.Application {
        const app = express();
        this.getMiddlewares().forEach((middleware) => {
            app.use(middleware);
        });

        app.post("/session", async (req: express.Request, res: express.Response) => {
            try {
                const result = await this.doCreateSession(req);
                res.status(200).json(result);
            } catch (error) {
                log.error("Error creating session on behalf of IAM", error, { error });
                // we treat all errors as bad request here and forward the error message to the caller
                res.status(400).json({ error, message: error.message });
            }
        });

        return app;
    }

    protected async doCreateSession(req: express.Request) {
        if (!OIDCCreateSessionPayload.is(req.body)) {
            throw new Error("Unexpected payload.");
        }
        const payload = req.body;
        OIDCCreateSessionPayload.validate(payload);
        const claims = payload.claims;

        const existingUser = await this.userService.findUserForLogin({
            candidate: {
                authId: claims.sub,
                authProviderId: claims.iss,
                authName: claims.name,
            },
        });
        const user = existingUser || (await this.createNewOIDCUser(payload));

        await new Promise<void>((resolve, reject) => {
            req.login(user, (err) => {
                if (err) {
                    reject(err);
                } else {
                    resolve();
                }
            });
        });
        return {
            sessionId: req.sessionID,
            userId: user.id,
        };
    }

    protected async createNewOIDCUser(payload: OIDCCreateSessionPayload): Promise<User> {
        const { claims, organizationId } = payload;

        // Until we support SKIM (or any other means to sync accounts) we create new users here as a side-effect of the login
        const user = await this.userService.createUser({
            identity: {
                authId: claims.sub,
                authProviderId: claims.iss,
                authName: claims.name,
                primaryEmail: claims.email,
            },
            userUpdate: (user) => {
                user.name = claims.name;
                user.avatarUrl = claims.picture;
                user.organizationId = organizationId;
            },
        });

        // Check: Is this the first login? If yes, make the logged-in user the owner
        const memberInfos = await this.teamDb.findMembersByTeam(organizationId);
        const firstMember = memberInfos.filter((m) => m.userId !== BUILTIN_INSTLLATION_ADMIN_USER_ID).length === 0;
        await this.teamDb.addMemberToTeam(user.id, organizationId);
        if (firstMember) {
            await this.teamDb.setTeamMemberRole(user.id, organizationId, "owner");
        }

        return user;
    }
}
