import { Code2, Sparkles, Play } from 'lucide-react';
import CodeBlock from '../components/CodeBlock';
import { useState } from 'react';

export default function Examples() {
  const [selectedCategory, setSelectedCategory] = useState('basics');

  const examples = {
    basics: [
      {
        title: 'Hello World',
        description: 'The classic first program in Bhasa',
        code: `লেখ("নমস্কার বিশ্ব!");
লেখ("Hello World!")`
      },
      {
        title: 'Variables and Math',
        description: 'Variable declarations and arithmetic operations',
        code: `ধরি x = ১০;
ধরি y = ২০;
ধরি যোগফল = x + y;
লেখ("যোগফল: " + যোগফল);

ধরি গুণফল = x * y;
লেখ("গুণফল: " + গুণফল);`
      },
      {
        title: 'Bengali Variable Names',
        description: 'Using Bengali identifiers in your code',
        code: `ধরি নাম = "রহিম";
ধরি বয়স = ২৫;
ধরি যোগফল = ৫০০০০;
ধরি শহর = "ঢাকা";

লেখ(নাম + " এর বয়স " + বয়স);
লেখ("তিনি " + শহর + " তে থাকেন");
লেখ("বেতন: " + যোগফল + " টাকা");`
      },
      {
        title: 'Boolean Logic',
        description: 'Working with boolean values and conditions',
        code: `ধরি সক্রিয় = সত্য;
ধরি নিষ্ক্রিয় = মিথ্যা;

যদি (সক্রিয় == সত্য) {
    লেখ("ব্যবহারকারী সক্রিয়");
}

যদি (!নিষ্ক্রিয়) {
    লেখ("সেবা চালু আছে");
}`
      }
    ],
    functions: [
      {
        title: 'Simple Function',
        description: 'Defining and calling functions',
        code: `ধরি যোগ = ফাংশন(a, b) {
    ফেরত a + b;
};

ধরি ফলাফল = যোগ(৫, ৩);
লেখ("যোগফল: " + ফলাফল);  // ৮`
      },
      {
        title: 'Fibonacci (Recursion)',
        description: 'Recursive function to calculate Fibonacci numbers',
        code: `ধরি fibonacci = ফাংশন(n) {
    যদি (n < ২) {
        ফেরত n;
    } নাহলে {
        ফেরত fibonacci(n - ১) + fibonacci(n - ২);
    }
};

লেখ("Fibonacci(০): " + fibonacci(০));  // ০
লেখ("Fibonacci(১): " + fibonacci(১));  // ১
লেখ("Fibonacci(৫): " + fibonacci(৫));  // ৫
লেখ("Fibonacci(১০): " + fibonacci(১০));  // ৫৫`
      },
      {
        title: 'Factorial',
        description: 'Calculating factorial using recursion',
        code: `ধরি factorial = ফাংশন(n) {
    যদি (n <= ১) {
        ফেরত ১;
    }
    ফেরত n * factorial(n - ১);
};

লেখ("৫! = " + factorial(৫));   // ১২০
লেখ("৭! = " + factorial(৭));   // ৫০৪০`
      },
      {
        title: 'Higher-Order Functions',
        description: 'Functions that return other functions',
        code: `ধরি makeMultiplier = ফাংশন(factor) {
    ফেরত ফাংশন(num) {
        ফেরত num * factor;
    };
};

ধরি double = makeMultiplier(২);
ধরি triple = makeMultiplier(৩);

লেখ(double(৫));   // ১০
লেখ(triple(৫));   // ১৫`
      }
    ],
    closures: [
      {
        title: 'Counter Closure',
        description: 'Creating a counter with encapsulated state',
        code: `ধরি makeCounter = ফাংশন() {
    ধরি count = ০;

    ফেরত ফাংশন() {
        count = count + ১;
        ফেরত count;
    };
};

ধরি counter = makeCounter();
লেখ(counter());  // ১
লেখ(counter());  // ২
লেখ(counter());  // ৩`
      },
      {
        title: 'Bank Account',
        description: 'Simulating a bank account with closures',
        code: `ধরি createAccount = ফাংশন(initialBalance) {
    ধরি balance = initialBalance;

    ফেরত ফাংশন(operation, amount) {
        যদি (operation == "deposit") {
            balance = balance + amount;
            ফেরত balance;
        } নাহলে {
            যদি (balance >= amount) {
                balance = balance - amount;
                ফেরত balance;
            }
            ফেরত -১;
        }
    };
};

ধরি account = createAccount(১০০০);
লেখ("জমা: " + account("deposit", ৫০০));     // ১৫০০
লেখ("উত্তোলন: " + account("withdraw", ২০০));  // ১৩০০`
      }
    ],
    arrays: [
      {
        title: 'Array Operations',
        description: 'Working with arrays and built-in functions',
        code: `ধরি সংখ্যা = [১, ২, ৩, ৪, ৫];

লেখ("দৈর্ঘ্য: " + দৈর্ঘ্য(সংখ্যা));    // ৫
লেখ("প্রথম: " + প্রথম(সংখ্যা));      // ১
লেখ("শেষ: " + শেষ(সংখ্যা));        // ৫

ধরি বাকিগুলো = বাকি(সংখ্যা);
লেখ(বাকিগুলো);  // [২, ৩, ৪, ৫]

ধরি নতুন = যোগ(সংখ্যা, ৬);
লেখ(নতুন);  // [১, ২, ৩, ৪, ৫, ৬]`
      },
      {
        title: 'Sum Array Elements',
        description: 'Calculate the sum of array elements',
        code: `ধরি সংখ্যা = [১০, ২০, ৩০, ৪০, ৫০];
ধরি যোগফল = ০;
ধরি i = ০;

যতক্ষণ (i < দৈর্ঘ্য(সংখ্যা)) {
    যোগফল = যোগফল + সংখ্যা[i];
    i = i + ১;
}

লেখ("যোগফল: " + যোগফল);  // ১৫০`
      },
      {
        title: 'Nested Arrays',
        description: 'Multi-dimensional arrays',
        code: `ধরি ম্যাট্রিক্স = [
    [১, ২, ৩],
    [৪, ৫, ৬],
    [৭, ৮, ৯]
];

লেখ(ম্যাট্রিক্স[০][০]);  // ১
লেখ(ম্যাট্রিক্স[১][১]);  // ৫
লেখ(ম্যাট্রিক্স[২][২]);  // ৯`
      }
    ],
    hashes: [
      {
        title: 'Hash Maps',
        description: 'Creating and accessing hash maps',
        code: `ধরি ব্যক্তি = {
    "নাম": "করিম",
    "বয়স": ২৫,
    "শহর": "ঢাকা",
    "সক্রিয়": সত্য
};

লেখ("নাম: " + ব্যক্তি["নাম"]);
লেখ("বয়স: " + ব্যক্তি["বয়স"]);
লেখ("শহর: " + ব্যক্তি["শহর"]);`
      },
      {
        title: 'Dynamic Hash Updates',
        description: 'Adding and modifying hash entries',
        code: `ধরি কনফিগ = {
    "হোস্ট": "localhost",
    "পোর্ট": ৮০৮০
};

// নতুন এন্ট্রি যোগ
কনফিগ["ডিবাগ"] = সত্য;
কনফিগ["টাইমআউট"] = ৩০০০;

লেখ(কনফিগ);`
      }
    ],
    strings: [
      {
        title: 'String Operations',
        description: 'Common string manipulation functions',
        code: `ধরি বার্তা = "নমস্কার বিশ্ব";

লেখ(উপরে(বার্তা));      // uppercase
লেখ(নিচে(বার্তা));       // lowercase
লেখ(দৈর্ঘ্য(বার্তা));    // length

ধরি পাঠ = "   স্পেস সহ   ";
লেখ(ছাঁটো(পাঠ));         // trim whitespace`
      },
      {
        title: 'String Split and Join',
        description: 'Splitting and joining strings',
        code: `ধরি বাক্য = "আমি বাংলায় গান গাই";
ধরি শব্দ = বিভক্ত(বাক্য, " ");
লেখ(শব্দ);  // ["আমি", "বাংলায়", "গান", "গাই"]

ধরি যুক্ত_বাক্য = যুক্ত(শব্দ, "-");
লেখ(যুক্ত_বাক্য);  // "আমি-বাংলায়-গান-গাই"`
      }
    ],
    controlflow: [
      {
        title: 'If-Else Statements',
        description: 'Conditional branching',
        code: `ধরি বয়স = ২০;

যদি (বয়স >= ১৮) {
    লেখ("প্রাপ্তবয়স্ক");
} নাহলে {
    লেখ("নাবালক");
}

ধরি স্কোর = ৮৫;

যদি (স্কোর >= ৯০) {
    লেখ("গ্রেড: A+");
} নাহলে {
    যদি (স্কোর >= ৮০) {
        লেখ("গ্রেড: A");
    } নাহলে {
        যদি (স্কোর >= ৭০) {
            লেখ("গ্রেড: B");
        } নাহলে {
            লেখ("গ্রেড: C");
        }
    }
}`
      },
      {
        title: 'While Loops',
        description: 'Iterating with while loops',
        code: `ধরি i = ১;

যতক্ষণ (i <= ১০) {
    লেখ("সংখ্যা: " + i);
    i = i + ১;
}

// Count down
ধরি count = ৫;
যতক্ষণ (count > ০) {
    লেখ(count);
    count = count - ১;
}
লেখ("শুরু!");`
      },
      {
        title: 'Break and Continue',
        description: 'Loop control with break and continue',
        code: `ধরি i = ০;

যতক্ষণ (i < ১০) {
    i = i + ১;

    // বিজোড় সংখ্যা এড়িয়ে যাও
    যদি (i % ২ == ১) {
        চালিয়ে_যাও;
    }

    লেখ("জোড় সংখ্যা: " + i);

    // ৮ এ পৌঁছালে থামো
    যদি (i == ৮) {
        বিরতি;
    }
}`
      }
    ]
  };

  const categories = [
    { id: 'basics', label: 'Basics', icon: Code2 },
    { id: 'functions', label: 'Functions', icon: Sparkles },
    { id: 'closures', label: 'Closures', icon: Play },
    { id: 'arrays', label: 'Arrays', icon: Code2 },
    { id: 'hashes', label: 'Hash Maps', icon: Sparkles },
    { id: 'strings', label: 'Strings', icon: Play },
    { id: 'controlflow', label: 'Control Flow', icon: Code2 }
  ];

  return (
    <div className="flex flex-col min-h-screen">
      {/* Header */}
      <section className="py-16 bg-gradient-to-br from-blue-50 to-indigo-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center">
            <h1 className="text-5xl font-bold text-slate-900 mb-4">
              Code Examples
            </h1>
            <p className="text-xl text-slate-600 max-w-3xl mx-auto">
              Explore practical examples demonstrating Bhasa's features and capabilities
            </p>
          </div>
        </div>
      </section>

      {/* Content */}
      <section className="flex-grow py-12 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          {/* Category Tabs */}
          <div className="flex flex-wrap gap-2 mb-8 justify-center">
            {categories.map((category) => {
              const IconComponent = category.icon;
              return (
                <button
                  key={category.id}
                  onClick={() => setSelectedCategory(category.id)}
                  className={`flex items-center space-x-2 px-4 py-2 rounded-lg font-medium transition-all ${
                    selectedCategory === category.id
                      ? 'bg-blue-600 text-white shadow-md'
                      : 'bg-slate-100 text-slate-700 hover:bg-slate-200'
                  }`}
                >
                  <IconComponent className="h-4 w-4" />
                  <span>{category.label}</span>
                </button>
              );
            })}
          </div>

          {/* Examples Grid */}
          <div className="grid lg:grid-cols-2 gap-8">
            {examples[selectedCategory as keyof typeof examples].map((example, index) => (
              <div
                key={index}
                className="bg-white rounded-xl shadow-lg border border-slate-200 overflow-hidden hover:shadow-xl transition-all"
              >
                <div className="p-6 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-slate-200">
                  <h3 className="text-xl font-bold text-slate-900 mb-2">
                    {example.title}
                  </h3>
                  <p className="text-sm text-slate-600">{example.description}</p>
                </div>
                <div className="p-6">
                  <CodeBlock code={example.code} />
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="py-16 bg-gradient-to-r from-blue-600 to-indigo-600">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl font-bold text-white mb-4">
            Ready to Write Your Own Code?
          </h2>
          <p className="text-lg text-blue-100 mb-6">
            Download Bhasa and start building amazing things
          </p>
          <a
            href="https://github.com/Uttam-Mahata/bhasa"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center space-x-2 px-6 py-3 bg-white text-blue-600 rounded-lg font-medium shadow-lg hover:shadow-xl transition-all"
          >
            <Code2 className="h-5 w-5" />
            <span>View on GitHub</span>
          </a>
        </div>
      </section>
    </div>
  );
}
