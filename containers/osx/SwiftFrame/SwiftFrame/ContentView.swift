//
//  ContentView.swift
//  SwiftFrame
//
//  Created by Derek Anderson on 2/24/24.
//

import SwiftUI
import WebKit


struct ContentView: View {
    let webView = WebView()
    var body: some View {
        VStack {
            WebView().frame(maxWidth: .infinity, maxHeight: .infinity)
        }
    }
}

#Preview {
    ContentView()
}

struct WebView: NSViewRepresentable {
    let webView: WKWebView

    init() {
        webView = WKWebView(frame: .zero)
    }

    func makeNSView(context: Context) -> WKWebView {
        return webView
    }

    func updateNSView(_ nsView: WKWebView, context: Context) {
        let request = URLRequest(url: URL(string: "http://localhost:8080")!)
        webView.load(request)
    }
}
