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
const URL = process.env.SPORTY_BET_URL;

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

(async function sportyBet() {
    let timeStamp = new Date().toISOString().split('T')[0];

    let options = new chrome.Options();
    options.addArguments('--no-sandbox');
    options.addArguments('--disable-dev-shm-usage');
    // options.addArguments('--headless');

    let driver = await new Builder()
        .forBrowser('chrome')
        .setChromeOptions(options)
        .build();

    try {
        await driver.get(URL);

        await driver.manage().setTimeouts({ implicit: 5000 });

        let username = await driver.findElement(By.xpath("//input[@name='phone' and @placeholder='Mobile Number']"));
        await clearBox(username);
        await username.sendKeys(USERNAME);
        await driver.sleep(2000);

        let password = await driver.findElement(By.xpath("//input[@type='password' or @placeholder='Password']"));
        await clearBox(password);
        await password.sendKeys(PASSWORD);
        await driver.sleep(2000);

        let loginBtn = await driver.findElement(By.xpath("//button[@name='logIn' and .='Login']"));
        await loginBtn.click();
        await driver.sleep(5000);

        try {
            let closeButton = await driver.findElement(By.xpath("//div[@class='m-winning-wrapper']//i[@class='m-icon-close']"));
            await closeButton.click();
        } catch (error) {
            if (error.name !== 'NoSuchElementError') {
                throw error;
            }
        }

        let fLeagues = await driver.findElements(By.css('div.match-league'));
        console.log(`Total number of leagues: ${fLeagues.length}`);
        let matchClicked = 0;

        for (let lgPost = 1; lgPost <= fLeagues.length; lgPost++) {
            let leagueName = await driver.findElement(By.xpath(`(//div[@class='match-league'])[${lgPost}]/div[@class='league-title']/span[@class='text']`)).getText();
            console.log(`League ${lgPost}: ${leagueName}`);
            if (!checkSimulated(leagueName)) {
                let doubleChanceButton = await driver.findElement(By.xpath(`(//div[@class='match-league'])[${lgPost}]//div[contains(text(), 'Double Chance')]`));
                await driver.wait(until.elementIsVisible(doubleChanceButton), 10000);
                await doubleChanceButton.click();
                console.log(`Clicked Double Chance for league ${lgPost}`);

                let matchPerLeague = await driver.findElements(By.xpath(`(//div[@class='match-league'])[${lgPost}]//div[contains(@class,'match-row')]`));
                console.log(`Found ${matchPerLeague.length} matches in league ${lgPost}`);

                for (let mPost = 1; mPost <= matchPerLeague.length; mPost++) {
                    try {
                        let matchHdc = await driver.findElement(By.xpath(`(//div[@class='match-league'])[${lgPost}]//div[contains(@class, 'match-row')][${mPost}]//div[@class='m-outcome'][1]`));
                        let matchHdcOc = parseFloat(await matchHdc.getText());
                        console.log(`Match ${mPost} odds: ${matchHdcOc}`);
                        if (matchHdcOc < MAXIMUM_ODD) {
                            await matchHdc.click();
                            matchClicked += 1;
                        }
                    } catch (error) {
                        console.log(`Error finding match odds for match ${mPost} in league ${lgPost}: ${error}`);
                        if (error.name !== 'NoSuchElementError') {
                            throw error;
                        }
                    }
                }
            }
            if (matchClicked > MAXIMUM_GAME) {
                break;
            }
        }

        console.log(`Total matches clicked: ${matchClicked}.`);

        let betslip = await driver.findElement(By.css("input[placeholder='min. 10']"));
        await clearBox(betslip);
        await driver.sleep(5000);
        await betslip.sendKeys(STAKE_AMOUNT);

        let actions = driver.actions();

        
    } finally {
        await driver.quit();
    }
})();
