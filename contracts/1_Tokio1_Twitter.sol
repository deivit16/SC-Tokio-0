// SPDX-License-Identifier: MIT
pragma solidity >= 0.8.1 <0.9.0;

/**
 * @title TwitterTokio
 * @dev Store & get tweets using struct
 */
contract TwitterTokio {

    // Define the Tweet struct to represent individual tweets
    struct Tweet {
        string username;
        string tweet;
        uint timestamp;
    }

    // Mapping to store tweets by their IDs
    mapping(uint256 => Tweet) private tweets;

    // Variable to keep track of the total number of tweets
    uint256 public totalTweets;
    
    // Mapping to track the last read message index for each user
    mapping(address => uint256) private lastReadMessageIndex;
    
    // Event emitted when a new tweet is published
    event TweetPublished(string indexed username, uint256 indexed tweetId, string tweet, uint256 timestamp);

    function setTweet(string memory _tweetContent, string memory _username) public {
        // Validate the content of the message not exceed 300 characters
        require(bytes(_tweetContent).length <= 300, "Tweet content exceeds maximum length");
        
        // Create a new Tweet with the provided username, tweet content, and current timestamp
        tweets[totalTweets] = Tweet(_username, _tweetContent, block.timestamp);
        // Increment the total number of tweets
        totalTweets++;

        // Emit an event to signal the publication of the new tweet
        emit TweetPublished(_username, totalTweets - 1, _tweetContent, block.timestamp);
    }

   
    function getTweet(uint256 _startIndex) public view returns (Tweet[] memory) {
        // Validate the startIndex is not out of range (total of Messages in the struct Tweet)
        require(_startIndex <= totalTweets, "Start index out of range");

        // Calculate the number of unread tweets based on the start index
        uint256 numUnreadTweets = totalTweets - _startIndex;
        // Create a list and populate it with unread tweets
        Tweet[] memory tweetList = new Tweet[](numUnreadTweets);
        for (uint256 i = _startIndex; i < totalTweets; i++) {
            tweetList[i - _startIndex] = tweets[i];
        }
        
        return tweetList;
    }

    function unreadMessages() public returns (Tweet[] memory) {
        // Retrieve the index of the last read message for the current user
        uint256 startIndex = lastReadMessageIndex[msg.sender]; 
        // Set the end index to the total number of tweets
        uint256 endIndex = totalTweets; 
        // Update the last read message index for the current user to the end index
        lastReadMessageIndex[msg.sender] = endIndex;
        // Retrieve and return unread messages between the start and end index
        return getLastUnreadMessage(startIndex, endIndex); 
    }

    function getLastUnreadMessage(uint256 _startIndex, uint256 _endIndex) private view returns (Tweet[] memory) {
        //// Ensure that the range of indices is valid
        require(_startIndex < _endIndex && _endIndex <= totalTweets, "Invalid range");
        
        // Calculate the number of messages within the specified range
        uint256 numMessages = _endIndex - _startIndex;

        // Initialize an array to store the retrieved messages
        Tweet[] memory result = new Tweet[](numMessages);
        for (uint256 i = _startIndex; i < _endIndex; i++) {
            // Populate the array with messages from the specified range
            result[i - _startIndex] = tweets[i];
        }
        return result;
    }

}

