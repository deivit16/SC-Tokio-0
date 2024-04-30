const TwitterTokio = artifacts.require("./TwitterTokio.sol");

module.exports = function (deployer) {
    //deployer.deploy(TwitterTokio, "0xda15492755a5db63578710c442aC4E3251705a3F");
    deployer.deploy(TwitterTokio);
};
