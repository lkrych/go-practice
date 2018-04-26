import React, { Component } from 'react';

class Search extends Component {
  constructor(props) {
    super(props);
    this.state = {
     
      searchTerm: "",
      searching: false,
      searchResults: {
        addresses: [],
        jsonSearched: "",
        itemSearched: ""
      }
    };
   
    this.searchTerm = this.searchTerm.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.logout = this.logout.bind(this);
  }

  logout(){
    localStorage.removeItem('id_token');
    localStorage.removeItem('access_token');
    localStorage.removeItem('profile');
    location.reload();
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
      fetch(`http://localhost:3000/api/search/${this.state.environment}/${this.state.selectedPartner}/${this.state.searchTerm}`,{
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
        
    

  render() {
    let loader = <div></div>;
    if (this.state.searching) {
      loader = <div className="loader"></div>;
    } 

    let searchResults = <div></div>;
    if(this.state.searchResults["addresses"].length !== 0) {
      searchResults = 
      
        <ul>
          {this.state.searchResults["addresses"].map( address => (
            <li key={address}>{address}</li>
          ))}
        </ul>;
    }

    let searchHeader= <div></div>;
    if(this.state.searchResults["jsonSearched"] !== "" && this.state.searchResults["itemSearched"] !== "") {
      searchHeader = 
      <div>
        <h3>Found {this.state.searchResults["jsonSearched"]} match(es) out of {this.state.searchResults["itemSearched"]} objects using {this.state.searchResults["searchTerm"]} as the search term(s) on the path {this.state.searchResults["searchPath"]}</h3>
        <br/>
        <h4>It took {this.state.searchResults["time"]} seconds</h4>
        {searchResults}
      </div>;
    }
    
  
      return (
          <div className="results-container">

          <h1>Search rss feeds for podcast info</h1>
          <button className="pull-right"><a onClick={this.logout}>Log out</a></button>
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
          {searchHeader}
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