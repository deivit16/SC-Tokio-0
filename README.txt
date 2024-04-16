# TwitterTokio Smart Contract

## Description
TwitterTokio is a smart contract that allows storing and retrieving tweets on the blockchain using a tweet data structure. Users can publish new tweets and retrieve previous tweets using functions provided by the contract.

## Features
- Publish new tweets with a 300-character limit.
- Retrieve previous tweets using a start index.
- Maintain a record of the last tweet read by each user.

## Contract Structure
The contract consists of the following main parts:
- **Tweet Struct**: Defines the structure of a tweet, which includes the username, tweet content, and timestamp.
- **Tweet Mapping**: Associates each tweet with a unique identifier.
- **Last Read Message Index Mapping**: Maintains a record of the last tweet read by each user.
- **Events**: Emits an event when a new tweet is published.
- **Public Functions**:
  - `setTweet`: Allows users to publish new tweets.
  - `getTweet`: Allows users to retrieve previous tweets starting from a given index.
  - `getLastUnreadMessage`: Allows users to get the index of the last unread tweet.

## Usage
To interact with the contract:
1. Deploy the contract on a compatible blockchain.
2. Use the provided functions to publish new tweets or retrieve previous tweets as needed.

## Notes
- Ensure that the length of the tweet content does not exceed 300 characters when calling the `setTweet` function.
- Ensure that the start index in the `getTweet` function is within the range of available tweets.
- Note that the start index in the `getTweet` function should be less than or equal to the total number of tweets stored in the contract.

## Author
This contract was created by DMR - TOKIO SCHOOL Spain.


