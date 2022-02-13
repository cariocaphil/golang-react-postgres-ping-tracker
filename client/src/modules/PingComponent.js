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

  const getLocale = () => {
    return navigator.languages && navigator.languages.length
      ? navigator.languages[0]
      : navigator.language;
  }

  const locale = getLocale();

  const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;

  const formatDate = (dateItem) => {
    const newDate = new Date(dateItem);
    return newDate.toLocaleString(locale, { timeZone: timezone });
  };

  return (
    <>
      <h1>Ping {pingState}</h1>
      <p>Most recent pings: </p>
      <ul>
        {occurrences &&
          occurrences.map((item) => {
            return <li>{formatDate(item)}</li>;
          })}
      </ul>
    </>
  );
}

export default PingComponent;
