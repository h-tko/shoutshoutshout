var path = require('path');

module.exports = {
    entry: path.join(__dirname, 'static/javascript/app.js'),
    output: {
        path: path.join(__dirname, 'static/dist'),
        filename: 'bundle.js',
    },
    cache: true,
    module: {
        loaders: [
            {
                test: /\.css$/, loader: 'style-loader!css-loader'
            }
        ],
    },
    devtool: 'source-map',
};
