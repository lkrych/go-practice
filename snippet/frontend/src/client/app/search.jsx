import React, { Component } from 'react';

import AudioPlayer from './audio_player';

class Search extends Component {
  constructor(props) {
    super(props);
    this.state = {
     
      searchTerm: "",
      searching: false,
      searchResults: {},
      playingAudio: false,
      currentlyPlaying: null,
    };
   
    this.searchTerm = this.searchTerm.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.playAudio = this.playAudio.bind(this);
    this.removeAudioPlayer = this.removeAudioPlayer.bind(this);
  }

  handleChange(e) {
    this.setState({
      searchTerm: e.target.value
    });
  }

  searchTerm(e) {
    e.preventDefault();
    this.setState({
      searching: true
    }, function() {
      const headers = buildRequestHeaders();
      fetch(`http://localhost:8080/api/search/${this.state.searchTerm}`,{
        credentials: 'include',
        headers: headers
      }).then(
        response => response.json().then(
          json => {
            this.setState({
              searching: false,
              searchResults: json
            });
          }
        ));
    });
    
  }

  playAudio(audioURL){
    let newState = !this.state.playingAudio;
    this.setState({
      playingAudio: newState,
      currentlyPlaying: audioURL,
    });
  }

  removeAudioPlayer(){
    let newState = !this.state.playingAudio;
    this.setState({
      playingAudio: newState,
      currentlyPlaying: null,
    });
  }
        

  render() {
    let loader = <div></div>;
    if (this.state.searching) {
      loader = <div className="loader"></div>;
    } 
    let audioPlayer = <div></div>;
    if (this.state.playingAudio) {
      audioPlayer = <AudioPlayer audioURL={this.state.currentlyPlaying} removeAudioPlayer={this.removeAudioPlayer} />;
    }

    let results =  <div></div>;
    if (Object.keys(this.state.searchResults).length > 1) {
      let podcast = this.state.searchResults;
      results = 
        <div>
          <img style={{"height": "200px", "width": "200px"}} src={podcast.image.url} alt={`image for ` + podcast.title}/>
          {audioPlayer}
          <h2>{podcast.title}</h2>
          <p>{podcast.description}</p>
          <ul>
            {podcast.items.map((el, idx) => {
              return (
                <li key={idx}>
                  <a onClick={this.playAudio(el.enclosures[0].url)}>
                    <h4>{el.title}</h4>
                  </a>
                  <p>{el.description}</p>
                  <p>{new Date(el.published).toLocaleString()}</p>
                </li>
              );
            })}
          </ul>
        </div>;
    }
    
  
    return (
        <div className="results-container">

        <h1>Search rss feeds for podcast info</h1>
        <div className="search-container"> 
          <div>
            <form onSubmit={this.searchTerm} >
              <div className="form">
                <label htmlFor="text_search">Search term</label> 
                
                <input id="text_search" type="text" value={this.state.value} onChange={this.handleChange}/>
                <br/>
                <input type="submit" value="Submit"/>
              </div>
            </form>
          </div>
        </div>
        <h1>Results</h1>
        {loader}
        {results}
      </div>
    );
  }
}

const buildRequestHeaders = () => {
  const headers = new Headers();
  headers.append('Content-Type', 'application/json');
  headers.append('Authorization', `Bearer ${localStorage.id_token}`);
  return headers;
};

export default Search;