import { useContext, useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { useAppDispatch, useAppSelector } from "../app/hooks";
import { store } from "../app/store";
import { updateEmail, EmailConfig } from "../features/configurationSlice";
import { APIContext, updateConfig } from "../services/API";

function EmailConfiguration() {
  const [domain, setDomain] = useState(store.getState().configuration.email.domain);
  const [port, setPort] = useState(store.getState().configuration.email.port);
  const [username, setUsername] = useState(store.getState().configuration.email.username);
  const [password, setPassword] = useState(store.getState().configuration.email.password);
  const [tls, setTLS] = useState(store.getState().configuration.email.tls);
  const dispatch = useAppDispatch()

    const api = useContext(APIContext)

  useEffect(() => {
    dispatch(updateEmail({
        domain: domain,
        port: port,
        username: username,
        password: password,
        tls: tls,
    }))
  }, [domain, port, username, password, tls])

  return (
    <div className="dark:bg-gray-800 rounded-lg p-4">
      <h2 className="text-2xl font-semibold">Email</h2>
      <div className="my-4 flex flex-row justify-between flex-wrap">
        <div>
          <label htmlFor="domain">Domain</label>
          <br />
          <input
            id="domain"
            type="text"
            value={domain}
            onChange={(event) => setDomain(event.target.value)}
            className="rounded-lg mt-2 focus:outline-none shadow-xl p-2 dark:bg-gray-700"
            placeholder="Domain"
          ></input>
        </div>
        <div className="">
          <label htmlFor="port">Port</label>
          <br />
          <input
            id="port"
            type="number"
            value={port}
            onChange={(event) => setPort(parseInt(event.target.value))}
            className="rounded-lg mt-2 focus:outline-none shadow-xl p-2 dark:bg-gray-700"
            placeholder="port"
          ></input>
        </div>

        <div className="">
          <label htmlFor="username">Username</label>
          <br />
          <input
            id="username"
            type="text"
            onChange={(event) => setUsername(event.target.value)}
            className="rounded-lg mt-2 focus:outline-none shadow-xl p-2 dark:bg-gray-700"
            placeholder="Username"
          ></input>
        </div>
        <div className="">
          <label htmlFor="password">Password</label>
          <br />
          <input
            id="password"
            type="password"
            onChange={(event) => setPassword(event.target.value)}
            className="rounded-lg mt-2 focus:outline-none shadow-xl p-2 dark:bg-gray-700"
            placeholder="Password"
          ></input>
        </div>
        <div>
          <label htmlFor="tls">TLS</label>
          <br />
          <input
            id="tls"
            type="checkbox"
            checked={tls}
            onChange={() =>  setTLS(!tls)}
            className="rounded-lg mt-2 focus:outline-none shadow-xl p-2 dark:bg-gray-700"
            placeholder="Domain"
          ></input>
        </div>
      </div>
      <div className="flex justify-end">
          <button onClick={() => {
              updateConfig(api, store.getState().configuration)
          }} className="rounded-lg text-lg p-2 px-4 bg-yellow-900 hover:bg-yellow-800 transition-all duration-200 active:outline-none focus:outline-none">Save</button>
      </div>
    </div>
  );
}

function NotionConfiguration() {
    return <div><h2>Notion</h2></div>
}

function GoogleDriveConfiguration() {
    return <div>
        <h2>Google Drive</h2>
    </div>
}

function Configuration() {
  return (
    <div className="container md:mx-32 dark:text-gray-400">
      <h1 className="text-4xl font-bold p-2">Configuration</h1>
      <EmailConfiguration />
    </div>
  );
}

export default Configuration;
