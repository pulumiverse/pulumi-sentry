import * as sentry from "@pulumiverse/sentry";

export const org = sentry.SentryOrganization.get("org", "example");
