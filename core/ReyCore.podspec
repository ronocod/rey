Pod::Spec.new do |s|
  s.name                = 'ReyCore'
  s.version             = '1.0'
  s.summary             = 'ReySLOCore'
  s.authors             = {"Conor O'Donnell"=>"mrconor@gmail.com"}
  s.license             = { :type => "Internal" }
  s.requires_arc        = true
  s.homepage            = 'https://github.com/ronocod/rey/'
  s.source              = { :git => "git@github.com:ronocod/rey.git", :tag => s.version }
  s.platform            = :ios, '9.0'
  s.preserve_paths      = 'build/ios/Core.framework'
  s.vendored_frameworks = 'build/ios/Core.framework'
end
