var path = require('path');

var APP_DIR = path.resolve(__dirname, 'src/client/app');

var config = {
  entry: APP_DIR + '/entry.jsx',
  output: {
    path: path.resolve(__dirname, 'build'), //must be absolute path, __dirname is a constant in nodeJS, a reference to cwd
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.jsx?/,
        loader: 'babel-loader',
        query: {
          presets: ['env','react']
        }
      },
    ]
  },
  mode: "production"
};

module.exports = config;