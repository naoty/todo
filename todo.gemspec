require "./lib/todo"

Gem::Specification.new do |spec|
  spec.name = "todo"
  spec.version = Todo::VERSION
  spec.license = "MIT"
  spec.summary = "My task manager"
  spec.author = "Naoto Kaneko"
  spec.email = "naoty.k@gmail.com"
  spec.files = Dir["lib/**/*.rb"] + Dir["bin/*"]
  spec.executable = "todo"
  spec.homepage = "https://github.com/naoty/todo"

  spec.add_development_dependency "rspec", "~> 3.10"
  spec.add_development_dependency "simplecov", "~> 0.21"
  spec.add_development_dependency "standard", "~> 1.1"
end
