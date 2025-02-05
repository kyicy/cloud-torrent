package server

//see github.com/jpillora/scraper for config specification
//cloud-torrent uses "<id>-item" handlers
var defaultSearchConfig = []byte(`{
	"eztv": {
	  "name": "EZTV",
	  "url": "https://eztv.ag/search/{{query}}",
	  "list": "table tr.forum_header_border",
	  "result": {
		"name": "td:nth-child(2) a",
		"url": ["td:nth-child(2) a", "@href"],
		"magnet": ["td:nth-child(3) a:nth-child(1)", "@href"],
		"size": "td:nth-child(4)",
		"seeds": "td:nth-child(6)"
	  }
	},
	"1337x": {
	  "name": "1337X",
	  "url": "http://1337x.to/sort-search/{{query}}/seeders/desc/{{page:1}}/",
	  "list": ".box-info-detail table.table tr",
	  "result": {
		"name": [".coll-1 a:nth-child(2)"],
		"url": [".coll-1 a:nth-child(2)", "@href"],
		"seeds": ".coll-2",
		"peers": ".coll-3",
		"size": [".coll-4", "/([\\d\\.]+ [KMGT]?B)/"]
	  }
	},
	"1337x/item": {
	  "name": "1337X (Item)",
	  "url": "http://1337x.to{{item}}",
	  "result": {
		"magnet": ["a[href^='magnet:']", "@href"]
	  }
	},
	"abb": {
	  "name": "The Audiobook Bay",
	  "url": "http://audiobookbay.nl/page/{{page:1}}?s={{query}}",
	  "list": "#content > .post",
	  "result": {
		"name": ["div.postTitle > h2 > a"],
		"url": ["div.postTitle > h2 > a", "@href"],
		"size": ["div.postContent", "/File Size: ([\\d\\.]+ [KMGT]?B)/"]
	  }
	},
	"abb/item": {
	  "name": "The Audiobook Bay (Item)",
	  "url": "http://audiobookbay.nl{{item}}",
	  "result": {
		"infohash": "/td>([a-f0-9]+)</",
		"tracker": "table.torrent_info tr:nth-child(1) td:nth-child(2)"
	  }
	},
	"lt": {
	  "name": "limetorrents",
	  "url": "https://www.limetorrents.cc/search/all/{{query}}/seeds/{{page:1}}/",
	  "list": ".table2 tbody tr",
	  "result": {
		"name": [".tt-name a:nth-child(2)"],
		"url": [".tt-name a:nth-child(2)", "@href"],
		"size": "td:nth-child(3)",
		"seeds": "td:nth-child(4)",
		"peers": "td:nth-child(5)"
	  }
	},
	"lt/item": {
	  "name": "limetorrents (Item)",
	  "url": "https://www.limetorrents.cc{{item}}",
	  "result": {
		"magnet": ["a.csprite_dltorrent[href^=\"magnet:\"]", "@href"]
	  }
	},
	"tpb": {
	  "name": "The Pirate Bay",
	  "url": "https://thepiratebay.org/search.php?/q={{query}}",
	  "list": "#torrents > li.list-entry",
	  "result": {
		"name": ".item-name a",
		"path": [".item-name a", "@href"],
		"magnet": [
		  ".item-icons a",
		  "@href"
		],
		"size": ".item-size",
		"seeds": ".item-seed",
		"peers": ".item-leech"
	  }
	},
	"nyaa": {
	  "name": "Nyaa",
	  "url": "https://nyaa.si/?f=0&c=0_0&s=seeders&o=desc&q={{query}}&p={{page:1}}",
	  "list": "table.torrent-list > tbody > tr",
	  "result": {
		"name": ["td:nth-child(2) a:nth-child(2)", "@title"],
		"path": ["td:nth-child(2) a:nth-child(2)", "@href"],
		"magnet": ["a[href^='magnet:']", "@href"],
		"size": "td:nth-child(4)",
		"seeds": "td:nth-child(6)",
		"peers": "td:nth-child(7)"
	  }
	}
  }`)
