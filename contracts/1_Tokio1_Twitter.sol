// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

/**
 * @title TwitterTokio
 * @dev Store & get tweets using struct
 */
contract TwitterTokio {
    struct Tweet {
        string username;
        string tweet;
        uint timestamp;
    }

    mapping(uint256 => Tweet) private tweets;
    uint256 public totalTweets;
    
    mapping(address => uint256) private lastReadMessageIndex;


    event TweetPublished(string indexed username, uint256 indexed tweetId, string tweet, uint256 timestamp);

    function setTweet(string memory _tweetContent, string memory _username) public {
        
        require(bytes(_tweetContent).length <= 300, "Tweet content exceeds maximum length");
        
        // Create a new Tweet
        tweets[totalTweets] = Tweet(_username, _tweetContent, block.timestamp);
        totalTweets++;
        emit TweetPublished(_username, totalTweets - 1, _tweetContent, block.timestamp);
    }

   
    function getTweet(uint256 _startIndex) public view returns (Tweet[] memory) {
        require(_startIndex <= totalTweets, "Start index out of range");

        uint256 numUnreadTweets = totalTweets - _startIndex;

        Tweet[] memory tweetList = new Tweet[](numUnreadTweets);
        for (uint256 i = _startIndex; i < totalTweets; i++) {
            tweetList[i - _startIndex] = tweets[i];
        }
        
        return tweetList;
    }

    function unreadMessages() public returns (Tweet[] memory) {
        uint256 startIndex = lastReadMessageIndex[msg.sender];
        uint256 endIndex = totalTweets;
        lastReadMessageIndex[msg.sender] = endIndex;
        return getLastUnreadMessage(startIndex, endIndex);
    }

    function getLastUnreadMessage(uint256 _startIndex, uint256 _endIndex) private view returns (Tweet[] memory) {
        require(_startIndex < _endIndex && _endIndex <= totalTweets, "Invalid range");
        uint256 numMessages = _endIndex - _startIndex;
        Tweet[] memory result = new Tweet[](numMessages);
        for (uint256 i = _startIndex; i < _endIndex; i++) {
            result[i - _startIndex] = tweets[i];
        }
        return result;
    }

}

