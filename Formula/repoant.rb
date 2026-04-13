class Repoant < Formula
  desc "Delete and mass delete GitHub repositories interactively"
  homepage "https://github.com/AasishDairelSahayaGrinspan/RepoAnt"
  version "0.2"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/AasishDairelSahayaGrinspan/RepoAnt/releases/download/ver0.2/repoant_darwin_arm64.tar.gz"
      sha256 "REPLACE_WITH_DARWIN_ARM64_SHA256_FROM_CHECKSUMS"
    else
      url "https://github.com/AasishDairelSahayaGrinspan/RepoAnt/releases/download/ver0.2/repoant_darwin_amd64.tar.gz"
      sha256 "REPLACE_WITH_DARWIN_AMD64_SHA256_FROM_CHECKSUMS"
    end
  end

  on_linux do
    url "https://github.com/AasishDairelSahayaGrinspan/RepoAnt/releases/download/ver0.2/repoant_linux_amd64.tar.gz"
    sha256 "REPLACE_WITH_LINUX_AMD64_SHA256_FROM_CHECKSUMS"
  end

  def install
    bin.install "repoant"
  end

  test do
    system "#{bin}/repoant", "version"
  end
end