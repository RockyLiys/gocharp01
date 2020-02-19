// @ts-ignore
import * as React from "react";
// @ts-ignore
import * as ReactDOM from "react-dom";

import { Hello } from "./components/home/Hello";

ReactDOM.render(
    <Hello compiler="TypeScript" framework="React" />,
    document.getElementById("example")
);