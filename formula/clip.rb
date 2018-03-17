class Clip < Formula
  desc "A simple key-value store for clipboard"
  homepage "https://github.com/mingrammer/clip"
  url "https://github.com/mingrammer/clip/releases/download/v0.0.2/clip_0.0.2_darwin_amd64.tar.gz"
  version "0.0.2"
  sha256 "c9de3b649edbb1580802391f9ac54b1912589df99e2eb53d15cdc5cefad90a40"

  def install
    bin.install "clip"
  end

  test do
    
  end
end
