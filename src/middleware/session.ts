import session from "express-session";
import dotenv from "dotenv";

dotenv.config();
const sessionSecret: string = process.env.SESSION_SECRET || "default_secret";

// Create a session middleware
const sessionMiddleware = session({
    secret: sessionSecret,
    resave: false,                  // Don't save session if unmodified
    saveUninitialized: false        // Prevent saving empty session
});

export default sessionMiddleware;