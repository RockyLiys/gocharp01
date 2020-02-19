// @ts-ignore
import * as React from 'react';

export interface HelloProps { compiler: string; framework: string; }

// @ts-ignore
export const Hello = (props: HelloProps) => <h1>Hello from {props.compiler} and {props.framework}!</h1>;