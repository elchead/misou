import React, { useState } from "react";
import { Result } from "../app/store";

class ExampleContainer extends React.Component<{ content: string }> {
  /**
   * Called after mounting the component. Triggers initial update of
   * the iframe
   */
  componentDidMount() {
    this._updateIframe();
  }

  /**
   * Called each time the props changes. Triggers an update of the iframe to
   * pass the new content
   */
  componentDidUpdate() {
    this._updateIframe();
  }

  /**
   * Updates the iframes content and inserts stylesheets.
   * TODO: Currently stylesheets are just added for proof of concept. Implement
   * and algorithm which updates the stylesheets properly.
   */
  _updateIframe() {
    const iframe: any = this.refs.iframe;
    const document = iframe.contentDocument;
    const head = document.getElementsByTagName("head")[0];
    document.body.innerHTML = this.props.content;
  }

  /**
   * This component renders just and iframe
   */
  render() {
    return <iframe className="w-full h-full" ref="iframe" />;
  }
}

function HtmlResultCard({ result, extended, callback }: { result: Result, extended: boolean, callback: Function }) {

  return (
    <div
      className={`rounded-md p-4 transition-all duration-1000 delay-100 dark:bg-gray-700 m-4 dark:text-gray-100 ${
        extended ?  "h-screen w-screen" : "w-screen h-32"
      }`}
      style={{}}
      onClick={() => callback()}
    >
      <div className={`rounded-md p-4 dark:bg-gray-600 h-full transition-all duration-500 `} >
        {result.matches > 0 ? (
          <p>
            <span className="font-semibold">{result.matches}</span>
            {result.matches === 1 ? "match" : "matches"}
          </p>
        ) : null}
        <p>
          {result.matches == 0 ? "Match " : null}
          in <span className="font-semibold">{result.title}</span>
          {result.link && result.link.length > 0 ? (
            <a
              className="underline hover:font-semibold"
              href={result.link}
              rel="noopener noreferrer"
              target="_blank"
            ></a>
          ) : null} 
        </p>
        {
          !extended ? (
            <div onClick={() => callback()} className="transition-all duration-500">
              <a
                className="underline hover:font-semibold"
                title="Click to show E-Mail"
              >
                show e-mail
              </a>
            </div>
          ) : <ExampleContainer content={result.content}/>
        }
      </div>
    </div>
  );
}

export default HtmlResultCard;
