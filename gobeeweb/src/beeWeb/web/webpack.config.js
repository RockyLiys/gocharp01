'use strict'
module.exports = function (env, args) {
    console.log(env)
    return {
        mode: env.production ? 'production' : 'development',  // 可以在cli上添加 webpack --mode=development
        entry: "./src/index.tsx",
        output: {
            filename: "bundle.js",
            path: __dirname + "/dist"
        },
        stats: {
            env: true
        },
        // Enable sourcemaps for debugging webpack's output.
        devtool: env.production ? 'source-maps' : 'eval',

        resolve: {
            // Add '.ts' and '.tsx' as resolvable extensions.
            extensions: [".ts", ".tsx", ".js", ".json"]
        },
        devServer: {
            // contentBase: path.join(__dirname, 'dist'),
            // compress: true,
            inline: true,
            port: 9000,
            open: true,
            hot: true,
            useLocalIp: true,  // 是否用本地ip
            onListening: function(server) {
                const port = server.listeningApp.address().port;
                console.log('Listening on port:', port);
            },
            proxy: {
                '/api': 'http://localhost:9876'
            }
        },
        performance: {
            assetFilter: function(assetFilename) {
                return assetFilename.endsWith('.js');
            }
        },
        module: {
            rules: [
                // All files with a '.ts' or '.tsx' extension will be handled by 'awesome-typescript-loader'.
                { test: /\.tsx?$/, loader: "awesome-typescript-loader" },

                // All output '.js' files will have any sourcemaps re-processed by 'source-map-loader'.
                { enforce: "pre", test: /\.js$/, loader: "source-map-loader" }
            ]
        },

        // When importing a module whose path matches one of the following, just
        // assume a corresponding global variable exists and use that instead.
        // This is important because it allows us to avoid bundling all of our
        // dependencies, which allows browsers to cache those libraries between builds.
        externals: {
            "react": "React",
            "react-dom": "ReactDOM"
        },
        plugins: []
    }
};
