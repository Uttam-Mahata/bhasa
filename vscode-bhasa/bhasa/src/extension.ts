// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from 'vscode';

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {

	// Use the console to output diagnostic information (console.log) and errors (console.error)
	// This line of code will only be executed once when your extension is activated
	console.log('Bhasa Language Extension is now active!');

	// Register the hello world command
	const disposable = vscode.commands.registerCommand('bhasa.helloWorld', () => {
		vscode.window.showInformationMessage('নমস্কার! Welcome to Bhasa Programming Language!');
	});

	context.subscriptions.push(disposable);

	// Register completion provider for Bhasa keywords
	const completionProvider = vscode.languages.registerCompletionItemProvider(
		{ language: 'bhasa', scheme: 'file' },
		{
			provideCompletionItems(document: vscode.TextDocument, position: vscode.Position) {
				const completions: vscode.CompletionItem[] = [];

				// Bengali keywords
				const bengaliKeywords = [
					{ label: 'ধরি', detail: 'Variable declaration (let)', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'ফাংশন', detail: 'Function', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'যদি', detail: 'If condition', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'নাহলে', detail: 'Else', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'ফেরত', detail: 'Return', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'যতক্ষণ', detail: 'While loop', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'পর্যন্ত', detail: 'For loop', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'বিরতি', detail: 'Break', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'চালিয়ে_যাও', detail: 'Continue', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'সত্য', detail: 'True', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'মিথ্যা', detail: 'False', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'লেখ', detail: 'Print function', kind: vscode.CompletionItemKind.Function },
				];

				// English keywords
				const englishKeywords = [
					{ label: 'let', detail: 'Variable declaration', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'function', detail: 'Function', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'if', detail: 'If condition', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'else', detail: 'Else', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'return', detail: 'Return', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'while', detail: 'While loop', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'for', detail: 'For loop', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'break', detail: 'Break', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'continue', detail: 'Continue', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'true', detail: 'True', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'false', detail: 'False', kind: vscode.CompletionItemKind.Keyword },
					{ label: 'print', detail: 'Print function', kind: vscode.CompletionItemKind.Function },
				];

				return [...bengaliKeywords, ...englishKeywords].map(item => {
					const completion = new vscode.CompletionItem(item.label, item.kind);
					completion.detail = item.detail;
					return completion;
				});
			}
		}
	);

	context.subscriptions.push(completionProvider);

	// Register hover provider for Bhasa keywords
	const hoverProvider = vscode.languages.registerHoverProvider('bhasa', {
		provideHover(document, position, token) {
			const range = document.getWordRangeAtPosition(position);
			const word = document.getText(range);

			const keywordDescriptions: { [key: string]: string } = {
				'ধরি': 'Variable declaration keyword (equivalent to `let`)',
				'ফাংশন': 'Function declaration keyword',
				'যদি': 'If condition keyword',
				'নাহলে': 'Else keyword',
				'ফেরত': 'Return statement keyword',
				'যতক্ষণ': 'While loop keyword',
				'পর্যন্ত': 'For loop keyword',
				'বিরতি': 'Break statement keyword',
				'চালিয়ে_যাও': 'Continue statement keyword',
				'সত্য': 'Boolean true value',
				'মিথ্যা': 'Boolean false value',
				'লেখ': 'Print function - outputs to console',
				'let': 'Variable declaration keyword',
				'function': 'Function declaration keyword',
				'if': 'If condition keyword',
				'else': 'Else keyword',
				'return': 'Return statement keyword',
				'while': 'While loop keyword',
				'for': 'For loop keyword',
				'break': 'Break statement keyword',
				'continue': 'Continue statement keyword',
				'true': 'Boolean true value',
				'false': 'Boolean false value',
				'print': 'Print function - outputs to console'
			};

			if (keywordDescriptions[word]) {
				return new vscode.Hover(keywordDescriptions[word]);
			}
		}
	});

	context.subscriptions.push(hoverProvider);
}

// This method is called when your extension is deactivated
export function deactivate() {}
