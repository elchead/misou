export function getHighlightedText(text: string, highlight: string) {
    // Split on highlight term and include term into parts, ignore case
    const parts = text.split(new RegExp(`(${highlight})`, 'gi'));
    return <span> { parts.map((part, i) => 
        <span key={i} className={part.toLowerCase() === highlight.toLowerCase() ? "search-highlight" : "" }>
            { part }
        </span>)
    } </span>;
}