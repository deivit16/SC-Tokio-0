import Web3 from "web3";
import TwitterTokio from "./contracts/TwitterTokio.json"
import {useState, useEffect} from "react";
import logo from './logo.svg';
import './App.css';

function App() {
  const [state,setState]=useState({ web3:null, contract:null });
  const [newTweet, setNewTweet] = useState(""); // Define newTweet
  const handleTweetChange = (event) => { // Define handleTweetChange
    setNewTweet(event.target.value);
  };

  const [data, setData]=useState([]);
  const [totalTweets, setTotalTweets] = useState(0);

  useEffect(()=>{
    async function connectWeb3(){
      const provider = new Web3.providers.HttpProvider("HTTP://127.0.0.1:7545");
      const web3 = new Web3(provider);
      //console.log(web3)
      const networkId = await web3.eth.net.getId();
      const deployedNetwork = TwitterTokio.networks[networkId];
      //console.log(deployedNetwork.address);

      const contract = new web3.eth.Contract(
        TwitterTokio.abi, 
        deployedNetwork.address
      );

      const accounts = await web3.eth.getAccounts();
      const account = accounts[0];  // Obtener la primera cuenta como dirección del remitente
      //console.log(account);
      setState({web3:web3, contract:contract, account:account})
    }
    
    connectWeb3();
  }, []);

  async function readData() {
    const { contract } = state;
    if (contract) {
      const tweetsArray = await contract.methods.getTweet(0).call();

      const tweets = [];
      tweetsArray.forEach(tweet => {
        const timestampInSeconds = parseInt(tweet.timestamp); // Convertir el timestamp a segundos
        const timestampInMilliseconds = timestampInSeconds * 1000; // Convertir segundos a milisegundos
        const date = new Date(timestampInMilliseconds); // Crear un objeto Date con el timestamp en milisegundos
        const formattedDate = date.toLocaleDateString('es-ES', { day: '2-digit', month: '2-digit', year: 'numeric' }); // Formatear la fecha como dd/MM/yyyy
        const tweetContent = tweet.tweet;
        const username = tweet.username;
        tweets.push({
          tweet: tweetContent,
          username: username,
          timestamp: formattedDate
        });
      });
      setData(tweets);
      console.log(tweets);
    }
  }

  async function fetchTotalTweets() {
    const { contract } = state;
    if (contract) {
      const total = await contract.methods
        .totalTweets()
        .call();
      setTotalTweets(parseInt(total));
    }
  }

  useEffect(() => {
    readData();
    fetchTotalTweets();
  }, [state]);

  async function writeData(){
    console.log("Entra WriteData");
    const { contract, account } = state;
    const userName = "@TokioSchool_Eth";
  
    await contract.methods
      .setTweet(newTweet, userName)
      .send({from: account, gas: '1000000' });
  
    const newTweetObject = {tweet: newTweet, username: userName};
    setData([...data, newTweetObject]);
    setNewTweet("");
    fetchTotalTweets();
  }
 
  return (
  <div className="App">
    <div className="header">
      <h1>Twitter Example Ethereum</h1>
      <p>Total de Tweets: {totalTweets}</p>
    </div>
    <div className="content">
      <div className="">
        <textarea
          className="tweet-input"
          placeholder="Escribe tu tweet aquí..."
          value={newTweet}
          onChange={handleTweetChange}
        ></textarea>
        <button className="tweet-button" onClick={writeData}>Send Tweet</button>
      </div>
      <div className="tweet">
        {data.map((tweet, index) => (
          <div key={index} className="tweet-item">
            <p className="author"></p>&gt;{tweet.tweet}
            <p className="tweet-content"></p>{tweet.timestamp}

            <p className="tweet-content">{tweet.username}</p>
          </div>
        ))}
      </div>
    </div>
  </div>);
}
export default App;
