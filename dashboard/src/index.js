import handle from "/exceptionHandler";

import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import injectTapEventPlugin from "react-tap-event-plugin";

// eslint-disable-next-line import/extensions
import "typeface-roboto";
import "/static";

import polyfill from "/polyfill";

import createClient from "/apollo/client";

import createStore from "/store";
import reducer from "/reducer";
import registerServiceWorker from "/registerServiceWorker";

import AppRoot from "/components/AppRoot";
import ErrorBoundary from "/components/util/ErrorBoundary";

polyfill().then(() => {
  // Configure store
  const store = createStore(reducer, {});

  const client = createClient();

  // Renderer
  ReactDOM.render(
    <ErrorBoundary handle={handle}>
      <BrowserRouter>
        <AppRoot reduxStore={store} apolloClient={client} />
      </BrowserRouter>
    </ErrorBoundary>,
    document.getElementById("root"),
  );
});

// Register React Tap event plugin
injectTapEventPlugin();

// Register service workers
registerServiceWorker();
