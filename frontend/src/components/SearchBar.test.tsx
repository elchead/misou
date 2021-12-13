import React from "react";
import { render, fireEvent, waitFor, screen } from "../test/test-utils";
import { SearchBarPure } from "./SearchBar";
import Results from "./Results";
import WS from "jest-websocket-mock";
import ReactTestUtils, { act } from "react-dom/test-utils"; // ES6
import { WSURL } from "../services/API";

const sleep = (milliseconds: number) => {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
};

// jest.setTimeout(10000);
// test("update input value upon entry", async () => {
//   const mockServer = new WS(WSURL, {
//     jsonProtocol: true,
//   });
//   const comp = render(
//     <>
//       <SearchBar />
//       <Results />
//     </>
//   );
//   const input = comp.getByPlaceholderText("Type to search");
//   await act(async () => {
//     await mockServer.connected;
//     fireEvent.change(input, { target: { value: "naval" } });

//     mockServer.send({
//       action: "results",
//       data: {
//         title: "Mocks are great",
//         content: "testing",
//         provider: "mock",
//       },
//       query: "naval",
//     });
//     await sleep(450); // need to wait for input to send WS query request
//   });
//   expect(input.value).toBe("naval");
//   await expect(mockServer).toHaveReceivedMessages([
//     { action: "query", query: "naval" },
//   ]);

//   // expect client to send msg -> receive results in Results
//   const resultTitle = await comp.findByTestId("resultTitle");
//   expect(resultTitle.textContent).toBe("Mocks are great");
// });

const debounceDelay = 300 + 50; // ms (needs some buffer)
describe("when query is entered", () => {
  let loadResults: jest.Mock<any, any>;
  let clearResults: jest.Mock<any, any>;
  let input: HTMLElement;
  beforeEach(() => {
    loadResults = jest.fn().mockName("loadResults");
    clearResults = jest.fn().mockName("loadResults");
    const sut = render(
      <SearchBarPure loadResults={loadResults} clearResults={clearResults} />
    );
    input = sut.getByPlaceholderText("Type to search...");
  });
  it("calls loadResults after debounce delay", async () => {
    await act(async () => {
      await changeDebouncedInput(input, "naval");
    });
    expect(loadResults).toHaveBeenCalledWith("naval");
  });
  it("deletes old results when query is empty", async () => {
    await act(async () => {
      await changeDebouncedInput(input, "naval");
    });
    expect(clearResults).not.toHaveBeenCalled();
    changeInput(input, "");
    expect(clearResults).toHaveBeenCalled();
  });
});

async function changeDebouncedInput(
  input: HTMLElement,
  value: string
): Promise<unknown> {
  changeInput(input, value);
  return sleep(debounceDelay);
}

function changeInput(input: HTMLElement, value: string) {
  return fireEvent.change(input, { target: { value: value } });
}
