import passport from "passport";
import OAuth2Strategy, { VerifyCallback } from "passport-oauth2";
import dotenv from "dotenv";

dotenv.config();
// TODO: Check TypeScript issues with process.env being undefined
const clientID: string = process.env.OSU_CLIENT_ID || "default_id";
const clientSecret: string = process.env.OSU_CLIENT_SECRET || "default_secret";
const callbackURL: string = process.env.OSU_CALLBACK_URL || "default_url";

passport.use(
    new OAuth2Strategy(
        {
            authorizationURL: "https://osu.ppy.sh/oauth/authorize",
            tokenURL: "https://osu.ppy.sh/oauth/token",
            clientID: clientID,
            clientSecret: clientSecret,
            callbackURL: callbackURL,
            scope: ["identify", "public"],
        },
        async (accessToken: string, refreshToken: string, profile: any, done: VerifyCallback) => { 
            try {
                // TODO: Implement the accesstoken fetching https://osu.ppy.sh/api/v2/me
                // Authorization: Bearer <access_token> | Content-Type: application/json | Accept: application/json
                return done(null, profile);
            } catch (error) {
                return done(error);
            }
        }
    )
);

export default passport;
