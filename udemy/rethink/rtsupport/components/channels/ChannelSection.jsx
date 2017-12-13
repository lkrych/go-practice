import React, {Component} from 'react';
import PropTypes from 'prop-types';
import ChannelForm from './ChannelForm.jsx';
import ChannelList from './ChannelList.jsx';

class ChannelSection extends Component{
  render(){
    return(
      <div>
        <ChannelList 
          channels={this.props.channels}
          setChannel={this.props.setChannel}
        />
        <ChannelForm
          addChannel={this.props.addChannel} />
      </div>
    );
  }

}

ChannelSection.propTypes = {
  channels: PropTypes.array.isRequired,
  setChannel: PropTypes.func.isRequired,
  addChannel: PropTypes.func.isRequired
};

export default ChannelSection;