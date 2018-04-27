import React, { Component } from 'react';

class AudioPlayer extends Component {
  constructor(props){
    super(props);
    this.logStar = this.logStar.bind(this);
  }

  logStar(){
    let audioPlayer = document.getElementById('media-player');
    console.log("The current time in the play of this file is: ", audioPlayer.currentTime);
    console.log("The duration of this file is: ", audioPlayer.duration);
  }

  render(){
    return(
      <div className="media-player" style={{"display": "flex", "flexDirection": "row", "alignItems": "center"}}>
        <audio id="media-player" controls src={this.props.audioURL}  />
        <a className="media-control" style={{"cursor": "pointer"}} onClick={this.logStar}><i className="fas fa-star"></i></a>
        <a className="media-control" style={{"cursor": "pointer"}} onClick={this.props.removeAudioPlayer}><i className="fas fa-times-circle"></i></a>
      </div>
    );
  }
}

export default AudioPlayer;