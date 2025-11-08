import {
  Book,
  FileText,
  Code2,
  Zap,
  Shield,
  Layers,
  BookOpen,
  FileCode,
  Globe,
  Settings,
  Database,
  Terminal,
  ExternalLink
} from 'lucide-react';

export default function Docs() {
  const docSections = [
    {
      title: 'Getting Started',
      icon: BookOpen,
      color: 'from-blue-500 to-indigo-500',
      docs: [
        {
          name: 'README',
          description: 'Project overview, features, and quick examples',
          file: 'README.md',
          icon: Book
        },
        {
          name: 'Quick Start',
          description: '2-minute guide to get up and running',
          file: 'docs/QUICKSTART.md',
          icon: Zap
        },
        {
          name: 'Usage Guide',
          description: 'Complete language usage reference',
          file: 'docs/USAGE.md',
          icon: Terminal
        }
      ]
    },
    {
      title: 'Language Features',
      icon: Code2,
      color: 'from-purple-500 to-pink-500',
      docs: [
        {
          name: 'Features Overview',
          description: 'Technical details and feature list',
          file: 'docs/FEATURES.md',
          icon: FileCode
        },
        {
          name: 'Bengali Implementation',
          description: 'Bengali script support details',
          file: 'docs/BENGALI_IMPLEMENTATION.md',
          icon: Globe
        },
        {
          name: 'Reserved Keywords',
          description: 'List of all Bengali keywords',
          file: 'docs/RESERVED_BENGALI_KEYWORDS.txt',
          icon: Book
        }
      ]
    },
    {
      title: 'Architecture & Design',
      icon: Layers,
      color: 'from-green-500 to-emerald-500',
      docs: [
        {
          name: 'Architecture',
          description: 'Deep technical architecture overview',
          file: 'docs/ARCHITECTURE.md',
          icon: Layers
        },
        {
          name: 'Architecture Summary',
          description: 'Quick technical reference guide',
          file: 'docs/ARCHITECTURE_SUMMARY.md',
          icon: FileText
        },
        {
          name: 'Compiler Design',
          description: 'Compiler implementation details',
          file: 'docs/COMPILER.md',
          icon: Code2
        },
        {
          name: 'Summary',
          description: 'Complete project achievement summary',
          file: 'docs/SUMMARY.md',
          icon: Book
        }
      ]
    },
    {
      title: 'Advanced Topics',
      icon: Database,
      color: 'from-orange-500 to-red-500',
      docs: [
        {
          name: 'Self-Hosting',
          description: 'Compiler written in Bhasa itself',
          file: 'docs/SELF_HOSTING.md',
          icon: FileCode
        },
        {
          name: 'Compiler API',
          description: 'Self-hosted compiler API reference',
          file: 'docs/COMPILER_API.md',
          icon: Code2
        },
        {
          name: 'OOP Status',
          description: 'Object-oriented programming implementation',
          file: 'docs/OOP_STATUS.md',
          icon: Settings
        },
        {
          name: 'Security Features',
          description: 'Security implementation summary',
          file: 'docs/SECURITY_SUMMARY.md',
          icon: Shield
        }
      ]
    }
  ];

  const quickLinks = [
    {
      title: 'GitHub Repository',
      description: 'View source code and contribute',
      url: 'https://github.com/Uttam-Mahata/bhasa',
      icon: Code2,
      color: 'bg-slate-900'
    },
    {
      title: 'Report Issues',
      description: 'Found a bug? Let us know',
      url: 'https://github.com/Uttam-Mahata/bhasa/issues',
      icon: Shield,
      color: 'bg-red-600'
    },
    {
      title: 'Example Programs',
      description: 'Browse code examples',
      url: 'https://github.com/Uttam-Mahata/bhasa/tree/main/examples',
      icon: FileCode,
      color: 'bg-blue-600'
    }
  ];

  return (
    <div className="flex flex-col">
      {/* Header */}
      <section className="py-16 bg-gradient-to-br from-blue-50 to-indigo-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <Book className="h-16 w-16 text-blue-600 mx-auto mb-4" />
            <h1 className="text-5xl font-bold text-slate-900 mb-4">
              Documentation
            </h1>
            <p className="text-xl text-slate-600 max-w-3xl mx-auto">
              Comprehensive guides and references for the Bhasa programming language
            </p>
          </div>
        </div>
      </section>

      {/* Documentation Sections */}
      <section className="py-12 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="space-y-12">
            {docSections.map((section, sectionIndex) => {
              const SectionIcon = section.icon;
              return (
                <div key={sectionIndex}>
                  {/* Section Header */}
                  <div className="flex items-center space-x-3 mb-6">
                    <div className={`p-3 bg-gradient-to-br ${section.color} rounded-lg shadow-lg`}>
                      <SectionIcon className="h-6 w-6 text-white" />
                    </div>
                    <h2 className="text-3xl font-bold text-slate-900">{section.title}</h2>
                  </div>

                  {/* Documents Grid */}
                  <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {section.docs.map((doc, docIndex) => {
                      const DocIcon = doc.icon;
                      return (
                        <a
                          key={docIndex}
                          href={`https://github.com/Uttam-Mahata/bhasa/blob/main/${doc.file}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="group block p-6 bg-white rounded-xl shadow-md border border-slate-200 hover:shadow-xl hover:border-blue-300 transition-all transform hover:-translate-y-1"
                        >
                          <div className="flex items-start justify-between mb-3">
                            <DocIcon className="h-8 w-8 text-blue-600" />
                            <ExternalLink className="h-4 w-4 text-slate-400 group-hover:text-blue-600 transition-colors" />
                          </div>
                          <h3 className="text-lg font-bold text-slate-900 mb-2 group-hover:text-blue-600 transition-colors">
                            {doc.name}
                          </h3>
                          <p className="text-sm text-slate-600 leading-relaxed">
                            {doc.description}
                          </p>
                          <div className="mt-4 text-xs text-slate-500 font-mono">
                            {doc.file}
                          </div>
                        </a>
                      );
                    })}
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </section>

      {/* Quick Links */}
      <section className="py-12 bg-slate-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-slate-900 mb-8 text-center">
            Quick Links
          </h2>

          <div className="grid md:grid-cols-3 gap-6">
            {quickLinks.map((link, index) => {
              const LinkIcon = link.icon;
              return (
                <a
                  key={index}
                  href={link.url}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="group block p-8 bg-white rounded-xl shadow-lg border border-slate-200 hover:shadow-2xl transition-all transform hover:-translate-y-1"
                >
                  <div className={`inline-flex p-4 rounded-lg ${link.color} mb-4`}>
                    <LinkIcon className="h-8 w-8 text-white" />
                  </div>
                  <h3 className="text-xl font-bold text-slate-900 mb-2 group-hover:text-blue-600 transition-colors flex items-center space-x-2">
                    <span>{link.title}</span>
                    <ExternalLink className="h-4 w-4" />
                  </h3>
                  <p className="text-slate-600">{link.description}</p>
                </a>
              );
            })}
          </div>
        </div>
      </section>

      {/* Documentation Stats */}
      <section className="py-12 bg-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="bg-gradient-to-r from-blue-600 to-indigo-600 rounded-2xl shadow-2xl p-8 text-center text-white">
            <h2 className="text-3xl font-bold mb-4">
              Comprehensive Documentation
            </h2>
            <p className="text-lg text-blue-100 mb-8">
              Over 15 detailed documentation files covering every aspect of the language
            </p>
            <div className="grid grid-cols-3 gap-6">
              <div>
                <div className="text-4xl font-bold mb-1">15+</div>
                <div className="text-sm text-blue-200">Documentation Files</div>
              </div>
              <div>
                <div className="text-4xl font-bold mb-1">3,200+</div>
                <div className="text-sm text-blue-200">Lines of Go Code</div>
              </div>
              <div>
                <div className="text-4xl font-bold mb-1">2,500+</div>
                <div className="text-sm text-blue-200">Lines of Bhasa Code</div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Additional Resources */}
      <section className="py-12 bg-gradient-to-br from-slate-50 to-blue-50">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <FileText className="h-12 w-12 text-blue-600 mx-auto mb-4" />
          <h2 className="text-3xl font-bold text-slate-900 mb-4">
            Need Help?
          </h2>
          <p className="text-lg text-slate-600 mb-6">
            Start with the Quick Start guide for a fast introduction, or dive deep into the
            Architecture documentation for technical details.
          </p>
          <div className="flex flex-wrap justify-center gap-4">
            <a
              href="https://github.com/Uttam-Mahata/bhasa/blob/main/docs/QUICKSTART.md"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center space-x-2 px-6 py-3 bg-blue-600 text-white rounded-lg font-medium shadow-lg hover:shadow-xl transition-all"
            >
              <Zap className="h-5 w-5" />
              <span>Quick Start Guide</span>
            </a>
            <a
              href="https://github.com/Uttam-Mahata/bhasa"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center space-x-2 px-6 py-3 bg-white text-slate-700 rounded-lg font-medium shadow-lg hover:shadow-xl border border-slate-200 transition-all"
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
