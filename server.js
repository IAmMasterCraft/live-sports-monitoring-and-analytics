require('dotenv').config();
const { Builder, By, Key, until } = require('selenium-webdriver');
const chrome = require('selenium-webdriver/chrome');
const fs = require('fs');

// Constants
const MAXIMUM_ODD = 1.2;
const MAXIMUM_GAME = 5;
const STAKE_AMOUNT = 100;
const USERNAME = process.env.SPORTY_BET_USERNAME;
const PASSWORD = process.env.SPORTY_BET_PASSWORD;

// Helper Functions
async function clearBox(webElement) {
    await webElement.sendKeys(Key.CONTROL, "a");
    await webElement.sendKeys(Key.BACK_SPACE);
}

function checkSimulated(name) {
    return name.toLowerCase().includes('simulated');
}