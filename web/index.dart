/** 
 * This Source Code is licensed under the MIT license. If a copy of the
 * MIT-license was not distributed with this file, You can obtain one at:
 * http://opensource.org/licenses/mit-license.html. 
 *
 * @author: Sean Westfall
 * @license MIT
 * @copyright Sean Westfall, 2015
 *     
 * name: RealTime Tech News 
 * version: 0.0.0
 * description: get tech news in real time
 * homepage: fieldsofgoldfish.com
 */
@HtmlImport('web/template/element/realtime_news.html')
library realtimenews.app;

// Imports
import 'dart:convert';                  // for converting json
import 'dart:html';                     // dart's HTML lib
import 'dart:async';                    // async lib
import 'package:polymer/polymer.dart';  // polymer

// Polymer Specific
import 'package:core_elements/core_animated_pages.dart';
import 'package:core_elements/core_animated_pages/transitions/slide_from_right.dart';
import 'package:core_elements/core_icon.dart';
import 'package:core_elements/core_icon_button.dart';
import 'package:core_elements/core_menu.dart';
import 'package:core_elements/core_scaffold.dart';
import 'package:core_elements/core_toolbar.dart';
import 'package:core_elements/roboto.dart';
import 'package:paper_elements/paper_item.dart';

// app specific
import 'src/elements.dart';

class Page {
    final String name;
    final String path;
    final bool isDefault;
    const Page(this.name, this.path, {this.isDefault: false});

    String toString() => '$name';
}

// does not extend Observable, (does not use two-way databinding)
@CustomTag('realtime-news')
class RealtimeNews extends PolymerElement {
    RealtimeNews.created() : super.created();

    // The list of pages in our app
    final List<Page> pages = const [
        const Page('One', 'one', isDefault: true),
        const Page('Two', 'two'),
        const Page('Three', 'three'),
        const Page('Four', 'four'),
        const Page('Five', 'five')
    ];
    
    // The current route
    @observable var route;
    // The current selected [Page]
    @observable Page selectedPage;

    // The [Router] that controls the app.
    final Router router = new Router(useFragment: true);

    // Convenience getters that return the expected types to avoid casts.
    CoreA11yKeys get keys => $['keys'];
    CoreScaffold get scaffold => $['scaffold'];
    CoreAnimatedPages get corePages => $['pages'];
    CoreMenu get menu => $['menu'];
    BodyElement get body => document.body;

    domReady() {
        // Set up the routes for all the pages.
        for (var page in pages) {
            router.root.addRoute(
                name: page.name, path: page.path, defaultRoute: page.isDefault,
                enter: enterRoute);
        }
        router.listen();

        // Set up the number keys to send you to pages.
        int i = 0;
        var keysToAdd = pages.map((page) => ++i);
        keys.keys = '${keys.keys} ${keysToAdd.join(' ')}';
    }

    // Updates [selectedPage] and the current route whenever the route changes.
    void routeChanged() {
    if (route is! String) return;
        if (route.isEmpty) {
            selectedPage = pages.firstWhere((page) => page.isDefault);
        } else {
            selectedPage = pages.firstWhere((page) => page.path == route);
        }
        router.go(selectedPage.name, {});
    }

    // Updates [route] whenever we enter a new route.
    void enterRoute(RouteEvent e) {
        route = e.path;
    }

    // Handler for key events.
    void keyHandler(e) {
        var detail = new JsObject.fromBrowserObject(e)['detail'];

        switch (detail['key']) {
            case 'left':
            case 'up':
                corePages.selectPrevious(false);
                return;
            case 'right':
            case 'down':
                corePages.selectNext(false);
                return;
            case 'space':
                detail['shift'] ? corePages.selectPrevious(false)
                    : corePages.selectNext(false);
            return;
        }

        // Try to parse as a number if we didn't recognize it as something else.
        try {
            var num = int.parse(detail['key']);
            if (num <= pages.length) {
                route = pages[num - 1].path;
            }
            return;
        } catch(e) {}
    }

    void menuItemClicked(_) {
        scaffold.closeDrawer();
    }

    void cyclePages(Event e, detail, sender) {
        var event = new JsObject.fromBrowserObject(e);
        // Clicks on links should not cycle pages.
        if (event['target'].localName == 'a') {
            return;
        }

        event['shiftKey'] ?
            sender.selectPrevious(true) : sender.selectNext(true);
    }
}
