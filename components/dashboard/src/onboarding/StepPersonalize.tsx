/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { User } from "@gitpod/gitpod-protocol";
import { FC, useCallback, useState } from "react";
import SelectIDEComponent from "../components/SelectIDEComponent";
import { ThemeSelector } from "../components/ThemeSelector";
import { OnboardingStep } from "./OnboardingStep";

type Props = {
    user: User;
    onComplete(ide: string, useLatest: boolean): void;
};
export const StepPersonalize: FC<Props> = ({ user, onComplete }) => {
    const [ide, setIDE] = useState(user?.additionalData?.ideSettings?.defaultIde || "code");
    const [useLatest, setUseLatest] = useState(user?.additionalData?.ideSettings?.useLatestVersion ?? false);

    // This step doesn't save the ide selection yet (happens at the end), just passes them along
    const handleSubmitted = useCallback(() => {
        onComplete(ide, useLatest);
    }, [ide, onComplete, useLatest]);

    const isValid = !!ide;

    return (
        <OnboardingStep
            title="How are you going to use Gitpod?"
            subtitle="We will tailor your experience based on your preferences."
            isValid={isValid}
            onSubmit={handleSubmitted}
        >
            <h3>Choose an editor</h3>
            <p>You can change this later in your user preferences.</p>
            <SelectIDEComponent
                onSelectionChange={(ide, latest) => {
                    setIDE(ide);
                    setUseLatest(latest);
                }}
                selectedIdeOption={ide}
                useLatest={useLatest}
            />

            <ThemeSelector className="mt-4" />
        </OnboardingStep>
    );
};
