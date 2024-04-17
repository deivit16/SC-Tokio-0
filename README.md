TwitterTokio Smart Contract

This smart contract, titled TwitterTokio, enables users to store and retrieve tweets using a struct-based approach. Users can publish tweets with associated usernames and timestamps, and retrieve tweets starting from a specified index. Additionally, the contract provides functionality for users to retrieve unread messages since their last read.

## Author
This contract was created by DMR - TOKIO SCHOOL Spain.

## License
This smart contract is licensed under the MIT License.

## Features
- Publish Tweets: Users can publish tweets with a maximum length of 300 characters.
- Retrieve Tweets: Users can retrieve tweets starting from a specified index.
- Unread Messages: Users can retrieve unread messages since their last read.

## Functions

### `setTweet(string memory _tweetContent, string memory _username)`

This function allows users to publish a new tweet. It requires the tweet content and the username of the author. The tweet is stored along with the username and timestamp of publication.

### `getTweet(uint256 _startIndex)`

Users can retrieve tweets starting from a specified index using this function. It returns an array of Tweet structs containing the tweet content, username, and timestamp.

### `unreadMessages()`

This function enables users to retrieve unread messages since their last read. It returns an array of Tweet structs containing unread tweets.

## Structs

### `Tweet`

The `Tweet` struct represents an individual tweet and consists of the following fields:

- `username`: The username of the tweet author.
- `tweet`: The content of the tweet.
- `timestamp`: The timestamp indicating when the tweet was published.

## Events

### `TweetPublished(string indexed username, uint256 indexed tweetId, string tweet, uint256 timestamp)`

This event is emitted when a new tweet is published. It includes the username of the author, the tweet ID, the tweet content, and the timestamp of publication.