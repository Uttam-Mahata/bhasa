import { Link } from 'react-router-dom';
import {
  Zap,
  Globe,
  Code2,
  Layers,
  Rocket,
  BookOpen,
  ChevronRight,
  Play,
  Terminal,
  Box,
  Cpu
} from 'lucide-react';
import CodeBlock from '../components/CodeBlock';

export default function Home() {
  const helloWorldCode = `লেখ("নমস্কার বিশ্ব!");`;

  const fibonacciCode = `ধরি fibonacci = ফাংশন(n) {
    যদি (n < ২) {
        ফেরত n;
    } নাহলে {
        ফেরত fibonacci(n - ১) + fibonacci(n - ২);
    }
};

লেখ(fibonacci(১০));  // Output: 55`;

  const bengaliVarsCode = `ধরি নাম = "রাজ";
ধরি বয়স = ২৫;
ধরি যোগফল = ৫০০০০;
লেখ(নাম + " এর বয়স " + বয়স);`;

  const features = [
    {
      icon: Zap,
      title: '3-10x Faster',
      description: 'Learn how bytecode compilation with a stack-based VM achieves significant performance gains over interpreted execution.',
      color: 'from-yellow-500 to-orange-500'
    },
    {
      icon: Globe,
      title: 'Bengali Script Native',
      description: 'Understand UTF-8 handling and Unicode support by working with Bengali keywords and identifiers.',
      color: 'from-blue-500 to-indigo-500'
    },
    {
      icon: Layers,
      title: 'Complete Compiler Pipeline',
      description: 'Study the full compilation process: Lexer → Parser → Compiler → VM with 41+ opcodes.',
      color: 'from-green-500 to-emerald-500'
    },
    // Note: Self-hosting compiler is in development
    // {
    //   icon: FileCode,
    //   title: 'Self-Hosting Compiler',
    //   description: 'The compiler is written entirely in Bhasa itself! A unique achievement showcasing language maturity.',
    //   color: 'from-green-500 to-emerald-500'
    // },
    {
      icon: Rocket,
      title: 'Rich Feature Set',
      description: 'Explore implementing functions, closures, arrays, hashes, recursion, and 30+ built-in functions.',
      color: 'from-red-500 to-rose-500'
    },
    {
      icon: Code2,
      title: 'Clean Implementation',
      description: 'Well-structured codebase with clear separation of concerns, perfect for learning compiler design.',
      color: 'from-cyan-500 to-blue-500'
    }
  ];

  const stats = [
    { label: 'Performance Boost', value: '3-10x', icon: Cpu },
    { label: 'Bengali Keywords', value: '13', icon: Globe },
    { label: 'Built-in Functions', value: '30+', icon: Terminal },
    { label: 'VM Opcodes', value: '41+', icon: Box }
  ];

  return (
    <div className="flex flex-col">
      {/* Hero Section */}
      <section className="relative overflow-hidden py-20 lg:py-32">
        <div className="absolute inset-0 bg-gradient-to-br from-blue-50 via-indigo-50 to-green-50 opacity-70"></div>
        <div className="absolute inset-0 bg-grid-slate-100 [mask-image:linear-gradient(0deg,white,rgba(255,255,255,0.6))] bg-[size:32px_32px]"></div>

        <div className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid lg:grid-cols-2 gap-12 items-center">
            {/* Left: Text Content */}
            <div className="space-y-8">
              <div className="inline-flex items-center space-x-2 px-4 py-2 bg-blue-100 text-blue-700 rounded-full text-sm font-medium">
                <Rocket className="h-4 w-4" />
                <span>Compiled Bengali Programming Language</span>
              </div>

              <h1 className="text-5xl lg:text-6xl font-bold leading-tight">
                <span className="bg-gradient-to-r from-orange-500 via-orange-600 to-green-600 bg-clip-text text-transparent">
                  ভাষা
                </span>
                <br />
                <span className="text-slate-900">Program in Bengali</span>
              </h1>

              <p className="text-xl text-slate-600 leading-relaxed">
                A learning project demonstrating compiler design with native Bengali script support.
                Explore how compilers work through a modern bytecode compiler and VM implementation,
                achieving 3-10x performance over tree-walking interpreters.
              </p>

              <div className="flex flex-wrap gap-4">
                <Link
                  to="/getting-started"
                  className="group inline-flex items-center space-x-2 px-6 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 text-white rounded-lg font-medium shadow-lg hover:shadow-xl transition-all transform hover:-translate-y-0.5"
                >
                  <Play className="h-5 w-5" />
                  <span>Get Started</span>
                  <ChevronRight className="h-4 w-4 group-hover:translate-x-1 transition-transform" />
                </Link>

                <Link
                  to="/examples"
                  className="inline-flex items-center space-x-2 px-6 py-3 bg-white text-slate-700 rounded-lg font-medium shadow-md hover:shadow-lg border border-slate-200 transition-all transform hover:-translate-y-0.5"
                >
                  <BookOpen className="h-5 w-5" />
                  <span>View Examples</span>
                </Link>
              </div>

              {/* Stats */}
              <div className="grid grid-cols-2 md:grid-cols-4 gap-4 pt-8">
                {stats.map((stat) => {
                  const IconComponent = stat.icon;
                  return (
                    <div key={stat.label} className="text-center">
                      <div className="inline-flex items-center justify-center w-10 h-10 bg-blue-100 rounded-lg mb-2">
                        <IconComponent className="h-5 w-5 text-blue-600" />
                      </div>
                      <div className="text-2xl font-bold text-slate-900">{stat.value}</div>
                      <div className="text-xs text-slate-600">{stat.label}</div>
                    </div>
                  );
                })}
              </div>
            </div>

            {/* Right: Code Examples */}
            <div className="space-y-4">
              <CodeBlock
                code={helloWorldCode}
                title="Hello World in Bhasa"
              />
              <CodeBlock
                code={bengaliVarsCode}
                title="Bengali Variable Names"
              />
            </div>
          </div>
        </div>
      </section>

      {/* Features Grid */}
      <section className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-4xl font-bold text-slate-900 mb-4">
              Why Learn Compiler Design with ভাষা?
            </h2>
            <p className="text-xl text-slate-600 max-w-3xl mx-auto">
              A practical, hands-on project to understand modern compiler architecture
              and virtual machine implementation.
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {features.map((feature, index) => {
              const IconComponent = feature.icon;
              return (
                <div
                  key={index}
                  className="group relative p-6 bg-white rounded-xl shadow-md hover:shadow-xl border border-slate-200 transition-all transform hover:-translate-y-1"
                >
                  <div className={`inline-flex p-3 rounded-lg bg-gradient-to-br ${feature.color} shadow-lg mb-4`}>
                    <IconComponent className="h-6 w-6 text-white" />
                  </div>
                  <h3 className="text-xl font-bold text-slate-900 mb-2">
                    {feature.title}
                  </h3>
                  <p className="text-slate-600 leading-relaxed">
                    {feature.description}
                  </p>
                </div>
              );
            })}
          </div>
        </div>
      </section>

      {/* Live Example Section */}
      <section className="py-20 bg-gradient-to-br from-slate-50 to-blue-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-4xl font-bold text-slate-900 mb-4">
              See It In Action
            </h2>
            <p className="text-xl text-slate-600">
              Fibonacci sequence with recursion
            </p>
          </div>

          <div className="max-w-3xl mx-auto">
            <CodeBlock
              code={fibonacciCode}
              title="fibonacci.bhasa"
              showLineNumbers
            />
          </div>

          <div className="text-center mt-8">
            <Link
              to="/examples"
              className="inline-flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium"
            >
              <span>Explore more examples</span>
              <ChevronRight className="h-4 w-4" />
            </Link>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 bg-gradient-to-r from-blue-600 to-indigo-600">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-4xl font-bold text-white mb-4">
            Ready to Learn Compiler Design?
          </h2>
          <p className="text-xl text-blue-100 mb-8">
            Explore the codebase and understand how compilers work from the ground up
          </p>
          <div className="flex flex-wrap justify-center gap-4">
            <Link
              to="/getting-started"
              className="inline-flex items-center space-x-2 px-8 py-4 bg-white text-blue-600 rounded-lg font-medium shadow-lg hover:shadow-xl transition-all transform hover:-translate-y-0.5"
            >
              <Rocket className="h-5 w-5" />
              <span>Get Started Now</span>
            </Link>
            <a
              href="https://github.com/Uttam-Mahata/bhasa"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center space-x-2 px-8 py-4 bg-blue-700 text-white rounded-lg font-medium shadow-lg hover:shadow-xl border border-blue-500 transition-all transform hover:-translate-y-0.5"
            >
              <Code2 className="h-5 w-5" />
              <span>View on GitHub</span>
            </a>
          </div>
        </div>
      </section>
    </div>
  );
}
