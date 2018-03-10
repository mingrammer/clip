class Clip < Formula
  desc "A simple key-value store for clipboard"
  homepage "https://github.com/mingrammer/clip"
  url "https://github.com/mingrammer/clip/releases/download/v0.0.1/clip_0.0.1_darwin_amd64.tar.gz"
  version "0.0.1"
  sha256 "c501d929db4db9bde80de32076b7c96271d9861f678a881e8cc98704ce8cc182"

  def install
    bin.install "clip"
  end

  test do
    
  end
end
