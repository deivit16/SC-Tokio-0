// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

/**
 * @title TwitterTokio
 * @dev Store & get tweets using struct
 */


contract DecentralizedMessaging {
    struct TweetMessage {
        address sender;
        string content;
        uint256 timestamp;
    }
    
    TweetMessage[] public messages;
    mapping(address => uint256) public lastReadIndex;

    function writeMessage(string memory _content) public {
        require(bytes(_content).length <= 300, "Message exceeds maximum length");
        messages.push(TweetMessage(msg.sender, _content, block.timestamp));
    }

    function readMessages(uint256 _startIndex) public returns (TweetMessage[] memory) {
        require(_startIndex < messages.length, "Start index out of bounds");
        uint256 endIndex = messages.length;
        if (lastReadIndex[msg.sender] < endIndex) {
            lastReadIndex[msg.sender] = endIndex;
        }
        return getMessagesInRange(_startIndex, endIndex);
    }

    function unreadMessages() public returns (TweetMessage[] memory) {
        uint256 startIndex = lastReadIndex[msg.sender];
        uint256 endIndex = messages.length;
        lastReadIndex[msg.sender] = endIndex;
        return getMessagesInRange(startIndex, endIndex);
    }

    function getMessagesInRange(uint256 _startIndex, uint256 _endIndex) private view returns (TweetMessage[] memory) {
        require(_startIndex < _endIndex && _endIndex <= messages.length, "Invalid range");
        uint256 numMessages = _endIndex - _startIndex;
        TweetMessage[] memory result = new TweetMessage[](numMessages);
        for (uint256 i = _startIndex; i < _endIndex; i++) {
            result[i - _startIndex] = messages[i];
        }
        return result;
    }
}
