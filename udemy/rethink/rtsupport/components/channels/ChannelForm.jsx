import React, {Component} from 'react';
import ChannelList from './ChannelList.jsx';

class ChannelForm extends Component{
  onSubmit(e){
    e.preventDefault();
    const node = this.refs.channel;
    const channelName = node.value;
    this.props.addChannel(channelName);
    node.value = '';
  }
  render(){
    return(
      <form onsubmit={this.onSubmit.bind(this)}>
        <input
        type='text'
        ref='channel'
        />
      </form>
    );
  }
}

ChannelForm.proptypes = {
  addChannel: React.Proptypes.func.isRequired
};

export default ChannelForm;