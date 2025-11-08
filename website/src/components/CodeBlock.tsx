import { useEffect, useRef } from 'react';
import Prism from 'prismjs';
import 'prismjs/themes/prism-tomorrow.css';
import { Copy, Check } from 'lucide-react';
import { useState } from 'react';

// Define Bhasa language for Prism
Prism.languages.bhasa = {
  comment: /\/\/.*/,
  string: {
    pattern: /"(?:\\.|[^\\"\r\n])*"/,
    greedy: true,
  },
  keyword: /\b(?:ধরি|ফাংশন|যদি|নাহলে|ফেরত|সত্য|মিথ্যা|যতক্ষণ|পর্যন্ত|বিরতি|চালিয়ে_যাও|নাল|অন্তর্ভুক্ত|let|function|if|else|return|true|false|while|for|break|continue|null|import)\b/,
  builtin: /\b(?:লেখ|দৈর্ঘ্য|প্রথম|শেষ|বাকি|যোগ|বিভক্ত|যুক্ত|উপরে|নিচে|ছাঁটো|প্রতিস্থাপন|খুঁজুন|শক্তি|বর্গমূল|পরম|সর্বোচ্চ|সর্বনিম্ন|গোলাকার|ফাইল_পড়ো|ফাইল_লেখো|ফাইল_যোগ|ফাইল_আছে|বাইট|ছোট_সংখ্যা|পূর্ণসংখ্যা|দীর্ঘ_সংখ্যা|দশমিক|দশমিক_দ্বিগুণ|অক্ষর_রূপান্তর|অক্ষর|কোড|অক্ষর_থেকে_কোড|সংখ্যা|লেখা)\b/,
  boolean: /\b(?:সত্য|মিথ্যা|true|false)\b/,
  null: /\b(?:নাল|null)\b/,
  number: /\b(?:০|১|২|৩|৪|৫|৬|৭|৮|৯|\d)+\b/,
  operator: /[+\-*\/%=!<>&|^~]+/,
  punctuation: /[{}[\]();,.:]/,
};

interface CodeBlockProps {
  code: string;
  language?: string;
  title?: string;
  showLineNumbers?: boolean;
}

export default function CodeBlock({
  code,
  language = 'bhasa',
  title,
  showLineNumbers = false
}: CodeBlockProps) {
  const codeRef = useRef<HTMLElement>(null);
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    if (codeRef.current) {
      Prism.highlightElement(codeRef.current);
    }
  }, [code, language]);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="relative group rounded-xl overflow-hidden shadow-lg border border-slate-200 bg-[#2d2d2d]">
      {/* Header */}
      {title && (
        <div className="flex items-center justify-between px-4 py-2 bg-slate-800 border-b border-slate-700">
          <span className="text-sm font-medium text-slate-300">{title}</span>
          <button
            onClick={handleCopy}
            className="p-1.5 text-slate-400 hover:text-white hover:bg-slate-700 rounded transition-all"
            title="Copy code"
          >
            {copied ? <Check className="h-4 w-4" /> : <Copy className="h-4 w-4" />}
          </button>
        </div>
      )}

      {/* Copy button (when no title) */}
      {!title && (
        <button
          onClick={handleCopy}
          className="absolute top-3 right-3 p-2 text-slate-400 hover:text-white hover:bg-slate-700 rounded-lg transition-all opacity-0 group-hover:opacity-100 z-10"
          title="Copy code"
        >
          {copied ? <Check className="h-4 w-4" /> : <Copy className="h-4 w-4" />}
        </button>
      )}

      {/* Code */}
      <div className="overflow-x-auto">
        <pre className={`!m-0 !bg-transparent ${showLineNumbers ? 'line-numbers' : ''}`}>
          <code ref={codeRef} className={`language-${language}`}>
            {code}
          </code>
        </pre>
      </div>
    </div>
  );
}
