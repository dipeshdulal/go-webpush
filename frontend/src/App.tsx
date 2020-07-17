import React from "react";
import "./App.css";
import { useNotification } from "./useNotification";

function App() {
  const notification = useNotification();

  return (
    <div className="App">
      <div
        style={{
          margin: 10,
          textAlign: "left",
        }}
      >
        <h3>SERVICE WORKER</h3>
        <pre
          style={{
            width: "100vw",
            overflow: "auto",
            padding: 10,
          }}
        >
          {JSON.stringify(notification, null, 2)}
        </pre>
      </div>
      <pre>STATUS: LOADING</pre>
      <div
        style={{
          display: "flex",
          padding: 20,
        }}
      >
        <div
          style={{
            flex: 1,
            textAlign: "left",
          }}
        >
          <h3>Send Message</h3>
          <input type="text" placeholder="Message" />
          <button onClick={() => {}}>Send Message</button>
        </div>
        <div
          style={{
            flex: 1,
            textAlign: "left",
          }}
        >
          <h3>Server Response</h3>
          <pre style={{}}>OLA</pre>
        </div>
      </div>
    </div>
  );
}

export default App;
