const webpack = require('webpack')
const path = require("path")
const AntdDayjsWebpackPlugin = require('antd-dayjs-webpack-plugin');

config = (env, args) => {
    console.log(env);
    console.log(args);
    return {
        mode: 'development',  // 可以在cli上添加 webpack --mode=development   env.production ? 'production' :
        entry: "./src/index.tsx",
        output: {
            filename: "bundle.js",
            path: __dirname + "/dist"
        },
        stats: {
            env: true
        },
        // Enable sourcemaps for debugging webpack's output.
        devtool: 'eval',  //env.production ? 'source-maps' :

        resolve: {
            // Add '.ts' and '.tsx' as resolvable extensions.
            extensions: [".ts", ".tsx", ".js", ".json"]
        },
        devServer: {
            // contentBase: path.join(__dirname, 'dist'),
            // compress: true,
            inline: true,
            host: '0.0.0.0',
            port: 9000,
            open: true,
            hot: true,
            useLocalIp: true,  // 是否用本地ip
            onListening: function(server) {
                const port = server.listeningApp.address().port;
                console.log('Listening on port:', port);
            },
            disableHostCheck: false,
            proxy: [{
                context: ['/api/**', '/u/**'],
                target: 'http://localhost:9876',
                secure: false
            }]
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
        plugins: [
            // 设置环境变量信息
            new webpack.DefinePlugin({
                PRODUCTION: JSON.stringify(true),
                VERSION: JSON.stringify('5fa3b9'),
                BROWSER_SUPPORTS_HTML5: true,
                TWO: '1+1',
                'typeof window': JSON.stringify('object'),
                'process.env': {
                    NODE_ENV: JSON.stringify(process.env.NODE_ENV)
                }
            }),
            new AntdDayjsWebpackPlugin()
            // ["import", { "libraryName": "antd", "style": "css" }] // `style: true` 会加载 less 文件
        ]
    }
};

module.exports = config