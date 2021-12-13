import { Result } from "../app/store";
import DriveIcon from "../assets/drive.svg";
import MailIcon from "../assets/mail.svg";
import NotionIcon from "../assets/notion.svg";
import DiskIcon from "../assets/disk.svg";
import DefaultIcon from "../assets/default.svg";

interface ResultCardProps {
  result: Result;
}

function ResultCard({ result }: ResultCardProps) {
  return (
    <div className="rounded-md p-4 dark:bg-gray-700 m-4 dark:text-gray-100 text-center flex flex-wrap flex-row align-center self-start min-w-max max-w-lg">
      <div className="rounded-md p-4 dark:bg-gray-600 flex items-center justify-center">
        {result.matches > 0 ? (
          <p>
            <span className="font-semibold">{result.matches}</span>
            {result.matches === 1 ? "match" : "matches"}
          </p>
        ) : null}
        <div>
          {result.matches === 0 ? "Match " : null}
          in{" "}
          <span className="font-semibold" data-testid="resultTitle">
            {result.title}
          </span>
        </div>
        <div>
          {result.link && result.link.length > 0 ? (
            <a
              className="underline hover:font-semibold"
              href={result.link}
              rel="noopener noreferrer"
              target="_blank"
            >
              <img
                className="h-7 w-7 float-right pl-2 max-w-full"
                src={parseProvider(result.provider)}
                title={result.provider}
              />
            </a>
          ) : (
            <img
              className="h-7 w-7 float-right pl-2 max-w-full"
              src={parseProvider(result.provider)}
              title={result.provider}
            />
          )}
        </div>
        <div>
          {result.content}
        </div>
      </div>
    </div>
  );
}

function parseProvider(provider: string): string {
  switch (provider) {
    case "gdrive":
      return DriveIcon;
    case "filesystem":
      return DiskIcon;
    case "email":
      return MailIcon;
    case "notion":
      return NotionIcon;
    default:
      return DefaultIcon;
  }
}

export default ResultCard;
