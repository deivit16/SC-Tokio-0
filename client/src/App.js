import Web3 from "web3";
import TwitterTokio from "./contracts/TwitterTokio.json"
import {useState, useEffect} from "react";
import logo from './logo.svg';
import './App.css';

function App() {
  const [state,setState]=useState({ web3:null, contract:null });
  const [data, setData]=useState([]);
  const [totalTweets, setTotalTweets] = useState(0);

  const [newTweet, setNewTweet] = useState(""); // Define newTweet
  const handleTweetChange = (event) => { // Define handleTweetChange
    setNewTweet(event.target.value);
  };

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

  useEffect(() => {
    
    async function readData(){
      const { contract } = state;
      if(contract) {
        const tweet = await contract.methods
          .getTweet(0)
          .call();
        setData([tweet]);
      }
    }
    readData();
  },[state]);

  useEffect(() => {
    async function fetchTotalTweets() {
      const { contract } = state;
      if (contract) {
        const total = await contract.methods
          .totalTweets()
          .call();
        setTotalTweets(total);
      }
    }
    fetchTotalTweets();
  }, [state]);

  async function writeData(){
    const { contract, account } = state;
    const userName = "@TokioSchool_Eth";

    // const time = Math.floor(Date.now() / 1000);
    // const data = document.querySelector("#value").value;
    await contract.methods
      .setTweet(userName, newTweet)
      .send({from: account, gas: '1000000' });
    setData([...data, {tweet: newTweet, username: userName}]);
    setNewTweet("")
    //window.location.reload();
  }

  return (
  <div className="App">
    <div className="header">
      <h1>Twitter Example Ethereum</h1>
      <p>Total de Tweets: {totalTweets}</p>
    </div>
    <div className="content">
      <div className="">
        {/* <input type="text" id="value"></input>
        <button onClick={writeData}>Send Tweet</button> */}
        <textarea
          className="tweet-input"
          placeholder="Escribe tu tweet aquí..."
          value={newTweet}
          onChange={handleTweetChange}
        ></textarea>
        <button className="tweet-button" onClick={writeData}>Enviar Tweet</button>
      </div>
      <div className="tweet">
        {data.map((tweet, index) => (
          <div key={index} className="tweet">
            <p className="author">{tweet.tweet}</p>
            <p>{tweet.username}</p>
          </div>
        ))}
      </div>
    </div>
  </div>);
}


export default App;
