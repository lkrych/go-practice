import React, { Component } from 'react';

class Search extends Component {
  constructor(props) {
    super(props);
    this.state = {
     
      searchTerm: "",
      searching: false,
      searchResults: {}
    };
   
    this.searchTerm = this.searchTerm.bind(this);
    this.handleChange = this.handleChange.bind(this);
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
        
    

  render() {
    let loader = <div></div>;
    if (this.state.searching) {
      loader = <div className="loader"></div>;
    } 

    let results =  <div></div>;
    if (Object.keys(this.state.searchResults).length > 1) {
      let podcast = this.state.searchResults;
      results = 
        <div>
          <img style={{"height": "200px", "width": "200px"}} src={podcast.image.url} alt={`image for ` + podcast.title}/>
          <h2>{podcast.title}</h2>
          <p>{podcast.description}</p>
          <ul>
            {podcast.items.map((el, idx) => {
              return (
                <li key={idx}>
                  <a href={el.enclosures[0].url}>
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