import React, { useState, useEffect } from "react";
import axios from "axios";

function PingComponent() {
  const [pingState, setPingState] = useState("");
  const [occurrences, setOccurrences] = useState("");

  useEffect(() => {
    axios
      .get("api/ping")
      .then((response) => {
        setPingState(response.data.message);
        setOccurrences(response.data.occurrences);
      })
      .catch(function (error) {
        console.log(error);
      });
  }, []);

  return (
    <>
      <h1>Ping {pingState}</h1>
      <p>Most recent pings: </p>
      <ul>
        {occurrences &&
          occurrences.map((item) => {
            return <li>{item}</li>;
          })}
      </ul>
    </>
  );
}

export default PingComponent;
