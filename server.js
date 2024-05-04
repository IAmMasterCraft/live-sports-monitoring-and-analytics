require('dotenv').config();
const { Builder, By, Key, until } = require('selenium-webdriver');
const chrome = require('selenium-webdriver/chrome');
const fs = require('fs');
const os = require('os');

// Constants
const MAXIMUM_ODD = 1.2;
const MAXIMUM_GAME = 5;
const STAKE_AMOUNT = 100;
const USERNAME = process.env.SPORTY_BET_USERNAME;
const PASSWORD = process.env.SPORTY_BET_PASSWORD;

const isMac = os.platform() === 'darwin';

// Helper Functions
async function clearBox(webElement) {
    const selectAll = isMac ? Key.chord(Key.COMMAND, 'a') : Key.chord(Key.CONTROL, 'a');
    await webElement.sendKeys(selectAll);
    await webElement.sendKeys(Key.BACK_SPACE);
}

function checkSimulated(name) {
    return name.toLowerCase().includes('simulated');
}