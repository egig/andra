const React = require('react');
const ReactDOM = require('react-dom/client');

import {
    HashRouter,
    Routes,
    Route,
} from "react-router-dom";

import { Provider } from 'react-redux'

import "normalize.css";
import "./index.css";

const App = () => {
    return <>[placehorlder]</>
}

const root = ReactDOM.createRoot(document.getElementById('root'));
// test
root.render(
    <Provider store={{}} >
        <HashRouter>
            <Routes>
                <Route path="/" element={<App />}>
                </Route>
            </Routes>
        </HashRouter>
    </Provider>
);
