import React, {Component} from 'react';
import PropTypes from 'prop-types';

class MessageList extends Component{
  render(){
    <ul>
      {this.props.messages.map( message => {
        return <Message key={message.id} message={message} />;
      })}
    </ul>;
  }
}

MessageList.propTypes = {
  messages: React.PropTypes.array.isRequired
};

export default MessageList;