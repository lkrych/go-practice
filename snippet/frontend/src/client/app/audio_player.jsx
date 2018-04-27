import React, { Component } from 'react';

class AudioPlayer extends Component {
  constructor(props){
    super(props);
    this.logStar = this.logStar.bind(this);
  }

  logStar(){
    let audioPlayer = document.getElementById('media-player');
    console.log("The current time in the play of this file is: ", audioPlayer.currentTime);
    console.log("The durationof this file is: ", audioPlayer.duration);

  }

  render(){
    return(
      <div className="media-player">
        <audio id="media-player" src={this.props.audioURL}  />
        <a onClick={this.logStar}><i class="fas fa-star"></i></a>
        <button onClick={this.props.removeAudioPlayer}><i class="fas fa-times-circle"></i></button>
      </div>
    );
  }
}

export default AudioPlayer;