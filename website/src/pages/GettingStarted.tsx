import {
  Download,
  Terminal,
  Play,
  FileCode,
  CheckCircle,
  Book,
  Github,
  Code2
} from 'lucide-react';
import CodeBlock from '../components/CodeBlock';

export default function GettingStarted() {
  const installSteps = [
    {
      number: 1,
      title: 'Clone the Repository',
      description: 'Get the latest version of Bhasa from GitHub',
      command: 'git clone https://github.com/Uttam-Mahata/bhasa.git\ncd bhasa',
      icon: Github
    },
    {
      number: 2,
      title: 'Build the Compiler',
      description: 'Build using Go directly or use Makefile for cross-platform binaries',
      command: '# Simple build for current platform\ngo build -o bhasa\n\n# OR using Makefile for optimized binary\nmake build\n\n# Build for all platforms (Linux, Windows, macOS)\nmake all\n\n# Build for specific platform\nmake linux-amd64    # Linux 64-bit\nmake windows-amd64  # Windows 64-bit\nmake darwin-arm64   # macOS Apple Silicon',
      icon: Terminal
    },
    {
      number: 3,
      title: 'Run Your First Program',
      description: 'Create and execute a Bhasa program',
      command: 'echo \'লেখ("নমস্কার বিশ্ব!");\' > hello.bhasa\n./bhasa hello.bhasa',
      icon: Play
    }
  ];

  const helloWorld = `লেখ("নমস্কার বিশ্ব!");`;

  const quickExample = `// Variables
ধরি নাম = "অনিকা";
ধরি বয়স = ২৫;

// Function
ধরি greet = ফাংশন(name) {
    ফেরত "নমস্কার, " + name + "!";
};

// Call function
লেখ(greet(নাম));

// Conditional
যদি (বয়স >= ১৮) {
    লেখ("প্রাপ্তবয়স্ক");
} নাহলে {
    লেখ("নাবালক");
}`;

  const replMode = `./bhasa
>> ধরি x = ১০;
>> ধরি y = ২০;
>> x + y
৩০
>> exit`;

  const requirements = [
    { name: 'Go', version: '1.19 or higher', description: 'Required to build the compiler' },
    { name: 'Git', version: 'Any recent version', description: 'To clone the repository' },
    { name: 'Terminal', version: 'Bash, Zsh, or equivalent', description: 'To run commands' }
  ];

  const features = [
    'Study lexer implementation with Bengali UTF-8 support',
    'Explore parser design with Pratt parsing',
    'Learn bytecode generation and optimization',
    'Understand VM execution with stack-based architecture',
    'Experiment with the interactive REPL',
    'Examine 30+ built-in function implementations'
  ];

  return (
    <div className="flex flex-col">
      {/* Header */}
      <section className="py-16 bg-gradient-to-br from-blue-50 to-indigo-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <Download className="h-16 w-16 text-blue-600 mx-auto mb-4" />
            <h1 className="text-5xl font-bold text-slate-900 mb-4">
              Getting Started
            </h1>
            <p className="text-xl text-slate-600 max-w-3xl mx-auto">
              Build and explore the Bhasa compiler to learn about language implementation
            </p>
          </div>
        </div>
      </section>

      {/* Requirements */}
      <section className="py-12 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-slate-900 mb-8 text-center">
            Requirements
          </h2>
          <div className="grid md:grid-cols-3 gap-6">
            {requirements.map((req, index) => (
              <div
                key={index}
                className="p-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-xl border border-blue-200"
              >
                <CheckCircle className="h-8 w-8 text-blue-600 mb-3" />
                <h3 className="text-lg font-bold text-slate-900 mb-1">{req.name}</h3>
                <p className="text-sm font-medium text-blue-600 mb-2">{req.version}</p>
                <p className="text-sm text-slate-600">{req.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Installation Steps */}
      <section className="py-12 bg-slate-50">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-slate-900 mb-8 text-center">
            Installation Steps
          </h2>

          <div className="space-y-8">
            {installSteps.map((step) => {
              const IconComponent = step.icon;
              return (
                <div key={step.number} className="relative">
                  <div className="flex items-start space-x-4">
                    {/* Step Number */}
                    <div className="flex-shrink-0 w-12 h-12 bg-gradient-to-br from-blue-600 to-indigo-600 rounded-full flex items-center justify-center shadow-lg">
                      <span className="text-white font-bold text-lg">{step.number}</span>
                    </div>

                    {/* Content */}
                    <div className="flex-grow">
                      <div className="bg-white rounded-xl shadow-md border border-slate-200 overflow-hidden">
                        <div className="p-6 border-b border-slate-200 bg-gradient-to-r from-blue-50 to-indigo-50">
                          <div className="flex items-center space-x-3 mb-2">
                            <IconComponent className="h-6 w-6 text-blue-600" />
                            <h3 className="text-xl font-bold text-slate-900">{step.title}</h3>
                          </div>
                          <p className="text-slate-600">{step.description}</p>
                        </div>
                        <div className="p-6">
                          <CodeBlock code={step.command} language="bash" />
                        </div>
                      </div>
                    </div>
                  </div>

                  {/* Connector Line */}
                  {step.number < installSteps.length && (
                    <div className="absolute left-6 top-12 w-0.5 h-8 bg-blue-300"></div>
                  )}
                </div>
              );
            })}
          </div>
        </div>
      </section>

      {/* Your First Program */}
      <section className="py-12 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-8">
            <FileCode className="h-12 w-12 text-indigo-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              Your First Bhasa Program
            </h2>
            <p className="text-lg text-slate-600">
              Create a file named <code className="px-2 py-1 bg-slate-100 rounded text-sm">hello.bhasa</code> with the following content
            </p>
          </div>

          <div className="mb-8">
            <CodeBlock code={helloWorld} title="hello.bhasa" />
          </div>

          <div className="bg-blue-50 border border-blue-200 rounded-xl p-6 mb-8">
            <h3 className="font-bold text-slate-900 mb-2 flex items-center space-x-2">
              <Terminal className="h-5 w-5 text-blue-600" />
              <span>Run Your Program</span>
            </h3>
            <CodeBlock code="./bhasa hello.bhasa" language="bash" />
          </div>

          <div className="bg-green-50 border border-green-200 rounded-xl p-6">
            <h3 className="font-bold text-slate-900 mb-2 flex items-center space-x-2">
              <CheckCircle className="h-5 w-5 text-green-600" />
              <span>Expected Output</span>
            </h3>
            <pre className="bg-slate-900 text-green-400 p-4 rounded-lg font-mono text-sm">
              নমস্কার বিশ্ব!
            </pre>
          </div>
        </div>
      </section>

      {/* Quick Example */}
      <section className="py-12 bg-slate-50">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-slate-900 mb-4 text-center">
            Quick Example
          </h2>
          <p className="text-center text-slate-600 mb-8">
            A more comprehensive example showcasing variables, functions, and conditionals
          </p>
          <CodeBlock code={quickExample} title="example.bhasa" showLineNumbers />
        </div>
      </section>

      {/* REPL Mode */}
      <section className="py-12 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-8">
            <Terminal className="h-12 w-12 text-green-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              Interactive REPL
            </h2>
            <p className="text-lg text-slate-600">
              Start the REPL for quick experimentation and testing
            </p>
          </div>

          <div className="mb-6">
            <CodeBlock code="./bhasa" language="bash" title="Start REPL" />
          </div>

          <CodeBlock code={replMode} language="bash" title="REPL Session Example" />
        </div>
      </section>

      {/* Cross-Platform Builds */}
      <section className="py-12 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-8">
            <Code2 className="h-12 w-12 text-blue-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              Cross-Platform Binary Builds
            </h2>
            <p className="text-lg text-slate-600">
              Build optimized binaries for multiple platforms using the included Makefile
            </p>
          </div>

          <div className="grid md:grid-cols-2 gap-6 mb-8">
            <div className="p-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-xl border border-blue-200">
              <h3 className="font-bold text-slate-900 mb-3 flex items-center space-x-2">
                <Terminal className="h-5 w-5 text-blue-600" />
                <span>Available Platforms</span>
              </h3>
              <ul className="space-y-2 text-sm text-slate-700">
                <li className="flex items-center space-x-2">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                  <span>Linux (AMD64, ARM64)</span>
                </li>
                <li className="flex items-center space-x-2">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                  <span>Windows (AMD64, ARM64)</span>
                </li>
                <li className="flex items-center space-x-2">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                  <span>macOS (Intel, Apple Silicon)</span>
                </li>
              </ul>
            </div>

            <div className="p-6 bg-gradient-to-br from-green-50 to-emerald-50 rounded-xl border border-green-200">
              <h3 className="font-bold text-slate-900 mb-3 flex items-center space-x-2">
                <Download className="h-5 w-5 text-green-600" />
                <span>Build Options</span>
              </h3>
              <ul className="space-y-2 text-sm text-slate-700">
                <li className="flex items-center space-x-2">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                  <span>Optimized with -ldflags for smaller binaries</span>
                </li>
                <li className="flex items-center space-x-2">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                  <span>All binaries output to bin/ directory</span>
                </li>
                <li className="flex items-center space-x-2">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                  <span>Run 'make help' for all available targets</span>
                </li>
              </ul>
            </div>
          </div>

          <div className="bg-slate-900 rounded-xl p-6">
            <h3 className="text-white font-bold mb-3">Quick Build Commands</h3>
            <CodeBlock
              code="# Build for all platforms\nmake all\n\n# Build for specific platforms\nmake linux        # All Linux binaries\nmake windows      # All Windows binaries\nmake darwin       # All macOS binaries\n\n# Clean build artifacts\nmake clean"
              language="bash"
            />
          </div>
        </div>
      </section>

      {/* What's Next */}
      <section className="py-12 bg-gradient-to-br from-slate-50 to-blue-50">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-8">
            <Book className="h-12 w-12 text-blue-600 mx-auto mb-4" />
            <h2 className="text-3xl font-bold text-slate-900 mb-4">
              What You Can Learn with Bhasa
            </h2>
          </div>

          <div className="grid md:grid-cols-2 gap-4 mb-8">
            {features.map((feature, index) => (
              <div
                key={index}
                className="flex items-start space-x-3 p-4 bg-white rounded-lg shadow-md border border-slate-200"
              >
                <CheckCircle className="h-6 w-6 text-green-600 flex-shrink-0 mt-0.5" />
                <span className="text-slate-700">{feature}</span>
              </div>
            ))}
          </div>

          <div className="bg-white rounded-xl shadow-lg border border-slate-200 p-8">
            <h3 className="text-2xl font-bold text-slate-900 mb-4 text-center">
              Next Steps
            </h3>
            <div className="grid md:grid-cols-3 gap-4">
              <a
                href="/examples"
                className="group p-6 bg-gradient-to-br from-blue-50 to-indigo-50 rounded-lg border border-blue-200 hover:shadow-lg transition-all text-center"
              >
                <Code2 className="h-10 w-10 text-blue-600 mx-auto mb-3" />
                <h4 className="font-bold text-slate-900 mb-2">Explore Examples</h4>
                <p className="text-sm text-slate-600">
                  Learn from practical code samples
                </p>
              </a>

              <a
                href="/features"
                className="group p-6 bg-gradient-to-br from-green-50 to-emerald-50 rounded-lg border border-green-200 hover:shadow-lg transition-all text-center"
              >
                <Book className="h-10 w-10 text-green-600 mx-auto mb-3" />
                <h4 className="font-bold text-slate-900 mb-2">Language Features</h4>
                <p className="text-sm text-slate-600">
                  Discover all Bhasa capabilities
                </p>
              </a>

              <a
                href="/docs"
                className="group p-6 bg-gradient-to-br from-green-50 to-emerald-50 rounded-lg border border-green-200 hover:shadow-lg transition-all text-center"
              >
                <FileCode className="h-10 w-10 text-green-600 mx-auto mb-3" />
                <h4 className="font-bold text-slate-900 mb-2">Documentation</h4>
                <p className="text-sm text-slate-600">
                  Read comprehensive guides
                </p>
              </a>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}
