import express from 'express';
import dotenv from 'dotenv';
import passport from 'passport';
import OAuth2Strategy from 'passport-oauth2';

import './strategies/osu-strategy';
import './middleware/session';
import sessionMiddleware from './middleware/session';

dotenv.config();

const app = express();

const PORT = process.env.PORT || 3000;

app.use(sessionMiddleware);

// Initialize passport
app.use(passport.initialize());
app.use(passport.session());

// TODO: Implement the routes

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
})