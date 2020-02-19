// @ts-ignore
import * as ReactDOM from "react-dom";

import React, { Component } from 'react';
import axios from "axios";
import { Button } from 'antd';

import { Hello } from "./components/home/Hello";

let service_conf = {
    baseURL: process.env.BASE_API, // api的base_url
    timeout: 5000 // request timeout
}
console.log(service_conf)
const service = axios.create(service_conf);

ReactDOM.render(
    <Hello compiler="TypeScript" framework="React" />,
    document.getElementById("example")
);

class App extends Component {
    render() {
        return (
            <div className="App">
                <Button type="primary">Button</Button>
            </div>
        );
    }
}

export default App;