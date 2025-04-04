CREATE TABLE short_urls (
    id BIGSERIAL PRIMARY KEY,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    creator_ip INET,  -- IP of the user who created the short link
    is_active BOOLEAN DEFAULT TRUE
);


CREATE TABLE url_clicks (
    id BIGSERIAL PRIMARY KEY,
    short_url_id BIGINT REFERENCES short_urls(id) ON DELETE CASCADE,
    clicked_at TIMESTAMP DEFAULT NOW(),
    ip_address INET,
    user_agent TEXT,
    referrer TEXT,
    country_code VARCHAR(5),
    region TEXT,
    city TEXT
);

CREATE INDEX idx_clicks_short_url_id ON url_clicks(short_url_id);
CREATE INDEX idx_clicks_clicked_at ON url_clicks(clicked_at);


CREATE TABLE url_clicks_aggregated (
    short_url_id BIGINT REFERENCES short_urls(id) ON DELETE CASCADE,
    click_date DATE NOT NULL,
    total_clicks BIGINT DEFAULT 0,
    PRIMARY KEY (short_url_id, click_date)
);


CREATE INDEX idx_short_urls_code ON short_urls(short_code);
CREATE INDEX idx_aggregated_clicks ON url_clicks_aggregated(short_url_id, click_date);
