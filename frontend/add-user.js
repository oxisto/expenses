#!/usr/bin/env node
const { hashSync } = require('bcrypt');
const { MongoClient } = require('mongodb');
const { question } = require('readline-sync');

const saltRounds = 10;

const url = 'mongodb://localhost:27017';
const dbName = 'expenses';
const collectionName = 'users';

async function upsertUser(username, password) {
    const client = await MongoClient.connect(url, { useNewUrlParser: true });

    const db = client.db(dbName);

    const users = db.collection(collectionName);

    const user = {
        username: username,
        password: hashSync(password, saltRounds)
    };

    await users.update({
        'username': username
    }, user, { upsert: true });

    console.log('Added or changed user ' + username + '.');

    client.close();
}

function main() {
    const username = question('Username: ');

    const password = question('Password: ', {
        hideEchoBack: true
    });

    upsertUser(username, password);
}

main();
