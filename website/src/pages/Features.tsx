import {
  Zap,
  Globe,
  Code2,
  FileCode,
  Box,
  Layers,
  Hash,
  Variable,
  Braces,
  GitBranch,
  Repeat,
  FunctionSquare,
  FileText,
  Calculator,
  Type,
  Download,
  CheckCircle
} from 'lucide-react';
import CodeBlock from '../components/CodeBlock';

export default function Features() {
  const coreFeatures = [
    {
      icon: Variable,
      title: 'Variables & Scoping',
      description: 'Lexical scoping with block-level variable declarations',
      example: `ধরি নাম = "রোহিত";
ধরি বয়স = ৩০;
ধরি সক্রিয় = সত্য;`
    },
    {
      icon: FunctionSquare,
      title: 'Functions & Closures',
      description: 'First-class functions with closure support and free variable capture',
      example: `ধরি makeCounter = ফাংশন() {
    ধরি count = ০;
    ফেরত ফাংশন() {
        count = count + ১;
        ফেরত count;
    };
};

ধরি counter = makeCounter();
লেখ(counter());  // 1
লেখ(counter());  // 2`
    },
    {
      icon: GitBranch,
      title: 'Control Flow',
      description: 'If-else conditionals, while loops, for loops with break and continue',
      example: `যদি (বয়স >= ১৮) {
    লেখ("প্রাপ্তবয়স্ক");
} নাহলে {
    লেখ("নাবালক");
}

যতক্ষণ (i < ১০) {
    লেখ(i);
    i = i + ১;
}`
    },
    {
      icon: Box,
      title: 'Arrays',
      description: 'Dynamic arrays with built-in methods for manipulation',
      example: `ধরি সংখ্যা = [১, ২, ৩, ৪, ৫];
লেখ(দৈর্ঘ্য(সংখ্যা));  // 5
লেখ(প্রথম(সংখ্যা));    // 1
লেখ(শেষ(সংখ্যা));     // 5

ধরি নতুন = যোগ(সংখ্যা, ৬);
লেখ(নতুন);  // [1, 2, 3, 4, 5, 6]`
    },
    {
      icon: Hash,
      title: 'Hash Maps',
      description: 'Key-value data structures with flexible types',
      example: `ধরি ব্যক্তি = {
    "নাম": "অর্জুন",
    "বয়স": ২৫,
    "শহর": "মুম্বাই"
};

লেখ(ব্যক্তি["নাম"]);  // অর্জুন
ব্যক্তি["ইমেইল"] = "arjun@example.com";`
    },
    {
      icon: Repeat,
      title: 'Recursion',
      description: 'Full support for recursive function calls',
      example: `ধরি factorial = ফাংশন(n) {
    যদি (n <= ১) {
        ফেরত ১;
    }
    ফেরত n * factorial(n - ১);
};

লেখ(factorial(৫));  // 120`
    }
  ];

  const builtinCategories = [
    {
      icon: FileText,
      title: 'I/O Functions',
      functions: ['লেখ() - Print output', 'ফাইল_পড়ো() - Read file', 'ফাইল_লেখো() - Write file', 'ফাইল_যোগ() - Append to file', 'ফাইল_আছে() - Check file exists']
    },
    {
      icon: Type,
      title: 'String Operations',
      functions: ['বিভক্ত() - Split', 'যুক্ত() - Join', 'উপরে() - Uppercase', 'নিচে() - Lowercase', 'ছাঁটো() - Trim', 'প্রতিস্থাপন() - Replace', 'খুঁজুন() - indexOf']
    },
    {
      icon: Box,
      title: 'Array Operations',
      functions: ['দৈর্ঘ্য() - Length', 'প্রথম() - First', 'শেষ() - Last', 'বাকি() - Rest', 'যোগ() - Push']
    },
    {
      icon: Calculator,
      title: 'Math Functions',
      functions: ['শক্তি() - Power', 'বর্গমূল() - Sqrt', 'পরম() - Abs', 'সর্বোচ্চ() - Max', 'সর্বনিম্ন() - Min', 'গোলাকার() - Round']
    }
  ];

  const bengaliKeywords = [
    { bengali: 'ধরি', english: 'let', description: 'Variable declaration' },
    { bengali: 'ফাংশন', english: 'function', description: 'Function declaration' },
    { bengali: 'যদি', english: 'if', description: 'Conditional' },
    { bengali: 'নাহলে', english: 'else', description: 'Else clause' },
    { bengali: 'ফেরত', english: 'return', description: 'Return statement' },
    { bengali: 'সত্য', english: 'true', description: 'Boolean true' },
    { bengali: 'মিথ্যা', english: 'false', description: 'Boolean false' },
    { bengali: 'যতক্ষণ', english: 'while', description: 'While loop' },
    { bengali: 'পর্যন্ত', english: 'for', description: 'For loop' },
    { bengali: 'বিরতি', english: 'break', description: 'Break statement' },
    { bengali: 'চালিয়ে_যাও', english: 'continue', description: 'Continue statement' },
    { bengali: 'নাল', english: 'null', description: 'Null value' },
    { bengali: 'অন্তর্ভুক্ত', english: 'import', description: 'Module import' }
  ];

  const architectureFeatures = [
    {
      icon: Code2,
      title: 'Lexer',
      description: 'UTF-8 tokenizer with Bengali character support and Bengali numerals (০-৯)',
      stats: '210 lines, 41 token types'
    },
    {
      icon: Layers,
      title: 'Parser',
      description: 'Pratt parser with 14 precedence levels for operator precedence',
      stats: '1,169 lines, AST generation'
    },
    {
      icon: FileCode,
      title: 'Compiler',
      description: 'Single-pass bytecode compiler with constant pool and jump patching',
      stats: '868 lines, 41+ opcodes'
    },
    {
      icon: Zap,
      title: 'Virtual Machine',
      description: 'Stack-based VM with 2048-element stack and 65,536 global slots',
      stats: '1,172 lines, 3-10x faster'
    }
  ];

  return (
    <div className="flex flex-col">
      {/* Header */}
      <section className="py-16 bg-gradient-to-br from-blue-50 to-indigo-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <h1 className="text-5xl font-bold text-slate-900 mb-4">
              Compiler Design Learning Features
            </h1>
            <p className="text-xl text-slate-600 max-w-3xl mx-auto">
              Explore compiler implementation through a complete feature set with Bengali script,
              bytecode compilation, and virtual machine execution
            </p>
          </div>
        </div>
      </section>

      {/* Bengali Keywords */}
      <section className="py-16 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <Globe className="h-12 w-12 text-blue-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              Bengali Keywords
            </h2>
            <p className="text-lg text-slate-600 max-w-2xl mx-auto">
              13 native Bengali keywords for writing expressive code
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
            {bengaliKeywords.map((keyword, index) => (
              <div
                key={index}
                className="p-4 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-lg border border-blue-200"
              >
                <div className="flex items-center justify-between mb-2">
                  <span className="text-2xl font-bold text-blue-600">{keyword.bengali}</span>
                  <span className="px-3 py-1 bg-white text-slate-700 rounded-md text-sm font-mono">
                    {keyword.english}
                  </span>
                </div>
                <p className="text-sm text-slate-600">{keyword.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Core Features */}
      <section className="py-16 bg-slate-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <Braces className="h-12 w-12 text-indigo-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              Implementation Examples
            </h2>
            <p className="text-lg text-slate-600">
              Learn how language features are implemented in a real compiler
            </p>
          </div>

          <div className="grid lg:grid-cols-2 gap-8">
            {coreFeatures.map((feature, index) => {
              const IconComponent = feature.icon;
              return (
                <div
                  key={index}
                  className="bg-white rounded-xl shadow-md border border-slate-200 overflow-hidden"
                >
                  <div className="p-6 border-b border-slate-200">
                    <div className="flex items-center space-x-3">
                      <div className="p-2 bg-indigo-100 rounded-lg">
                        <IconComponent className="h-6 w-6 text-indigo-600" />
                      </div>
                      <div>
                        <h3 className="text-xl font-bold text-slate-900">
                          {feature.title}
                        </h3>
                        <p className="text-sm text-slate-600">{feature.description}</p>
                      </div>
                    </div>
                  </div>
                  <div className="p-6 bg-slate-50">
                    <CodeBlock code={feature.example} />
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </section>

      {/* Built-in Functions */}
      <section className="py-16 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <Download className="h-12 w-12 text-green-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              Built-in Functions
            </h2>
            <p className="text-lg text-slate-600">
              30+ built-in functions for common operations
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
            {builtinCategories.map((category, index) => {
              const IconComponent = category.icon;
              return (
                <div
                  key={index}
                  className="p-6 bg-gradient-to-br from-green-50 to-emerald-50 rounded-xl border border-green-200"
                >
                  <div className="flex items-center space-x-2 mb-4">
                    <IconComponent className="h-6 w-6 text-green-600" />
                    <h3 className="text-lg font-bold text-slate-900">{category.title}</h3>
                  </div>
                  <ul className="space-y-2">
                    {category.functions.map((fn, idx) => (
                      <li key={idx} className="flex items-start space-x-2 text-sm text-slate-700">
                        <CheckCircle className="h-4 w-4 text-green-600 mt-0.5 flex-shrink-0" />
                        <span>{fn}</span>
                      </li>
                    ))}
                  </ul>
                </div>
              );
            })}
          </div>
        </div>
      </section>

      {/* Architecture */}
      <section className="py-16 bg-gradient-to-br from-slate-50 to-blue-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <Layers className="h-12 w-12 text-blue-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              Compilation Architecture
            </h2>
            <p className="text-lg text-slate-600">
              Professional-grade compiler with VM execution
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
            {architectureFeatures.map((feature, index) => {
              const IconComponent = feature.icon;
              return (
                <div
                  key={index}
                  className="relative p-6 bg-white rounded-xl shadow-lg border border-slate-200"
                >
                  <div className="absolute top-4 right-4 w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                    <span className="text-sm font-bold text-blue-600">{index + 1}</span>
                  </div>
                  <IconComponent className="h-10 w-10 text-blue-600 mb-4" />
                  <h3 className="text-xl font-bold text-slate-900 mb-2">
                    {feature.title}
                  </h3>
                  <p className="text-sm text-slate-600 mb-3">{feature.description}</p>
                  <div className="px-3 py-1 bg-blue-50 text-blue-700 rounded-md text-xs font-medium inline-block">
                    {feature.stats}
                  </div>
                </div>
              );
            })}
          </div>

          {/* Pipeline Diagram */}
          <div className="bg-white rounded-xl shadow-lg border border-slate-200 p-8">
            <h3 className="text-2xl font-bold text-slate-900 mb-6 text-center">
              Compilation Pipeline
            </h3>
            <div className="flex flex-col md:flex-row items-center justify-between space-y-4 md:space-y-0 md:space-x-4">
              {['Bengali Source', 'Lexer', 'Parser', 'Compiler', 'VM', 'Output'].map((stage, index) => (
                <div key={stage} className="flex items-center space-x-4">
                  <div className="flex flex-col items-center">
                    <div className="w-32 h-20 bg-gradient-to-br from-blue-600 to-indigo-600 rounded-lg shadow-lg flex items-center justify-center">
                      <span className="text-white font-bold text-sm text-center px-2">
                        {stage}
                      </span>
                    </div>
                  </div>
                  {index < 5 && (
                    <div className="hidden md:block text-slate-400">
                      <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                      </svg>
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Self-Hosting - Currently in development */}
      {/* <section className="py-16 bg-gradient-to-r from-blue-600 to-indigo-600">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <FileCode className="h-16 w-16 text-white mx-auto mb-6" />
          <h2 className="text-4xl font-bold text-white mb-4">
            Self-Hosting Compiler
          </h2>
          <p className="text-xl text-blue-100 mb-8">
            The Bhasa compiler is written entirely in Bhasa itself! This unique achievement
            demonstrates language maturity and showcases its capability to build complex systems.
          </p>
          <div className="inline-flex items-center space-x-2 px-6 py-3 bg-white text-blue-600 rounded-lg font-medium shadow-lg">
            <CheckCircle className="h-5 w-5" />
            <span>9 self-hosted modules, 2500+ lines of Bhasa code</span>
          </div>
        </div>
      </section> */}
    </div>
  );
}
