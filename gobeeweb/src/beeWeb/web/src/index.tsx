// @ts-ignore
import * as React from "react";
// @ts-ignore
import * as ReactDOM from "react-dom";
import axios from "axios";
import { Hello } from "./components/home/Hello";

let service_conf = {
    baseURL: process.env.BASE_API, // api的base_url
    timeout: 5000 // request timeout
}
const service = axios.create(service_conf);

ReactDOM.render(
    <Hello compiler="TypeScript" framework="React" />,
    document.getElementById("example")
);