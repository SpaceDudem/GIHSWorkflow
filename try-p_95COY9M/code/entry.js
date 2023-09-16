import { axios } from "@pipedream/platform";

export default defineComponent({
  props: {
    _csrfToken: {
      type: "string",
      label: "_csrfToken",
    },
    animal_ids: {
      type: "string",
      label: "Animal IDs",
    },
    apiKey: {
      type: "string",
      label: "API Key",
    },
    cookie: {
      type: "string",
      label: "Cookie",
    },
  },
  async run({ $ }) {
    const url = "https://www.shelterluv.com/api/v3/kennel_cards/get-shelter-template";
    const data = `_csrfToken=${this._csrfToken}&animal_ids=${this.animal_ids}`;

    const response = await axios($, {
      method: "POST",
      url: url,
      data: data,
      headers: {
        "sec-ch-ua": `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`,
        "Accept": "text/plain, */*; q=0.01",
        "Content-Type": "application/x-www-form-urlencoded",
        "X-Requested-With": "XMLHttpRequest",
        "sec-ch-ua-mobile": "?0",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
        "sec-ch-ua-platform": `"Windows"`,
        "Sec-Fetch-Site": "same-origin",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Dest": "empty",
        "host": "www.shelterluv.com",
        "X-Api-Key": this.apiKey,
        "Cookie": this.cookie,
      },
    });

    return { kennelCardUrl: response };
  },
});